package service

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"koriebruh/management/domain"
	"koriebruh/management/dto"
	"koriebruh/management/repository"
	"koriebruh/management/utils"
	"net/http"
)

type AuthService interface {
	Register(ctx *fiber.Ctx, request dto.RegisterRequest) (dto.WebResponse, error)
	Login(ctx *fiber.Ctx, request dto.LoginRequest) (uint, error)
	FindAllAdmin(ctx context.Context) (dto.WebResponse, error)
}

type AuthServiceImpl struct {
	*gorm.DB
	repository.AuthRepository
	*validator.Validate
}

func NewAuthService(DB *gorm.DB, authRepository repository.AuthRepository, validate *validator.Validate) *AuthServiceImpl {
	return &AuthServiceImpl{DB: DB, AuthRepository: authRepository, Validate: validate}
}

func (service AuthServiceImpl) Register(ctx *fiber.Ctx, request dto.RegisterRequest) (dto.WebResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return utils.ErrorResponseWeb(utils.ErrBadRequest, err), fmt.Errorf("there are those who do not meet the criteria")
	}

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		passwordHashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		request.Password = string(passwordHashed)

		registerData := domain.Admin{
			Username: request.Username,
			Password: request.Password,
			Email:    request.Email,
		}

		if err := service.AuthRepository.Register(ctx, tx, registerData); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return utils.ErrorResponseWeb(utils.ErrBadRequest, err), nil
	}

	return utils.SuccessRes(http.StatusCreated, "CREATED", "SUCCESS CREATED"), nil
}

func (service AuthServiceImpl) Login(ctx *fiber.Ctx, request dto.LoginRequest) (uint, error) {

	if err := service.Validate.Struct(request); err != nil {
		return 0, err
	}

	//FOP DAN SAVE INTO JWT
	var adminId uint

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		loginData := domain.Admin{
			Username: request.Username,
			Password: request.Password,
		}

		var err error
		adminId, err = service.AuthRepository.Login(ctx, tx, loginData)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return adminId, nil

}

func (service AuthServiceImpl) FindAllAdmin(ctx context.Context) (dto.WebResponse, error) {
	var admins []domain.Admin          // ini untuk tangkap hasil db
	var adminList []dto.AdminsResponse // ini dto

	err := service.DB.Transaction(func(tx *gorm.DB) error {

		admin, err := service.AuthRepository.FindAllAdmin(ctx, tx)
		if err != nil {
			return err
		}

		admins = admin
		return nil
	})

	if err != nil {
		return utils.ErrorResponseWeb(utils.ErrBadRequest, err), nil
	}

	for _, admin := range admins {
		adminLimit := dto.AdminsResponse{
			Username:  admin.Username,
			Email:     admin.Email,
			CreatedAt: admin.CreatedAt,
		}
		adminList = append(adminList, adminLimit)
	}

	return utils.SuccessRes(http.StatusOK, "OK", adminList), nil
}
