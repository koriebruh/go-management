package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"koriebruh/management/cnf"
	"koriebruh/management/dto"
	"koriebruh/management/service"
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
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}
	// Panggil service untuk menangani logika registrasi
	err := controller.AuthService.Register(ctx, request)
	if err != nil {
		// Menangani error yang terjadi
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	// Respons sukses
	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]string{
			"message": "success created new admin",
		},
	})
}

func (controller AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	// Parse request body ke DTO
	var request dto.LoginRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}
	// Panggil service untuk menangani logika registrasi
	adminId, err := controller.AuthService.Login(ctx, request)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
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
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   "ERROR Generate Token",
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  "token",
		Path:  "/",
		Value: tokenValue,
	})

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]string{
			"message": "Login Success",
			"token":   tokenValue,
		},
	})

}

func (controller AuthControllerImpl) Logout(ctx *fiber.Ctx) error {
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
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "StatusInternalServerError",
			Data:   nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   admins,
	})
}
