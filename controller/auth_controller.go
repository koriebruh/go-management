package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"koriebruh/management/cnf"
	"koriebruh/management/dto"
	"koriebruh/management/service"
	"koriebruh/management/utils"
	"net/http"
	"time"
)

type AuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	FindAllAdmin(ctx *fiber.Ctx) error
}

type AuthControllerImpl struct {
	service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{AuthService: authService}
}

func (controller AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	// Parse request body ke DTO
	var request dto.RegisterRequest
	if err := ctx.BodyParser(&request); err != nil {
		errResponse := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	webResponse, err := controller.AuthService.Register(ctx, request)
	if err != nil {
		errResponse := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

func (controller AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	// Parse request body ke DTO
	var request dto.LoginRequest
	if err := ctx.BodyParser(&request); err != nil {
		errResponse := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// Panggil service untuk menangani logika registrasi
	adminId, err := controller.AuthService.Login(ctx, request)
	if err != nil {
		errResponse := utils.ErrorResponseWeb(utils.ErrBadRequest, err)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	//SETTING GENERATE JWT
	expTime := time.Now().Add(time.Hour * 1) // << KADALUARSA DALAM 1 JAM
	claims := cnf.JWTClaim{
		UserID:   int(adminId),
		UserName: request.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "koriebruh",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenValue, err := tokenAlgo.SignedString([]byte(cnf.JWT_KEY))
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  "token",
		Path:  "/",
		Value: tokenValue,
	})

	res := utils.SuccessRes(http.StatusOK, "SUCCESS", map[string]string{"token": tokenValue})
	return ctx.Status(http.StatusOK).JSON(res)

}

func (controller AuthControllerImpl) Logout(ctx *fiber.Ctx) error {

	// VALIDATE TOKEN FROM HEADER
	token := ctx.Get("Authorization")
	if token == "" || len(token) < 7 || token[:7] != "Bearer " {
		errResponse := utils.ErrorResponseWeb(utils.ErrUnauthorized, fmt.Errorf("error token not found"))
		return ctx.Status(http.StatusUnauthorized).JSON(errResponse)
	}

	//DELETE TOKEN IN COOKIE CLIENT
	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour * 12),
		HTTPOnly: true,
	})

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]string{
			"message": "LogOut Success",
		},
	})
}

func (controller AuthControllerImpl) FindAllAdmin(ctx *fiber.Ctx) error {
	admins, err := controller.AuthService.FindAllAdmin(ctx.Context())
	if err != nil {
		webErr := utils.ErrorResponseWeb(utils.ErrInternalServerError, err)
		return ctx.Status(http.StatusInternalServerError).JSON(webErr)
	}

	return ctx.Status(http.StatusOK).JSON(admins)
}
