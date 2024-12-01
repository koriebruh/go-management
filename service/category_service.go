package service

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"koriebruh/management/cnf"
	"koriebruh/management/domain"
	"koriebruh/management/dto"
	"koriebruh/management/repository"
	"log"
)

type CategoryService interface {
	Create(ctx *fiber.Ctx, request dto.CategoryRequest) error
	FindAllCategory(ctx *fiber.Ctx) ([]dto.CategoryResponse, error)
	SummaryCategory(ctx context.Context) ([]dto.SummaryCategory, error)
}

type CategoryServiceImpl struct {
	*gorm.DB
	repository.CategoryRepository
	*validator.Validate
}

func NewCategoryService(DB *gorm.DB, categoryRepository repository.CategoryRepository, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{DB: DB, CategoryRepository: categoryRepository, Validate: validate}
}

func (service CategoryServiceImpl) Create(ctx *fiber.Ctx, request dto.CategoryRequest) error {
	if err := service.Validate.Struct(request); err != nil {
		return err
	}
	return service.DB.Transaction(func(tx *gorm.DB) error {
		//TAKE TOKEN
		token := ctx.Cookies("token")

		claims := &cnf.JWTClaim{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cnf.JWT_KEY), nil
		})

		//VERIFICATION JWT
		if err != nil || !tkn.Valid {
			errors.New("UNAUTHORIZED")
		}
		id := claims.UserID

		// MAPPING DTO TO ENTITY

		category := domain.Category{
			Name:        request.Name,
			Description: request.Description,
			CreatedBy:   uint(id),
		}

		err = service.CategoryRepository.Create(ctx.Context(), tx, &category)
		if err != nil {
			log.Print("Error di repository:", err)
			return err
		}

		return nil
	})
}

func (service CategoryServiceImpl) FindAllCategory(ctx *fiber.Ctx) ([]dto.CategoryResponse, error) {
	var categories []domain.Category

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		category, err := service.CategoryRepository.FindAllCategory(ctx.Context(), tx)
		if err != nil {
			return err
		}

		categories = category
		return nil
	})

	if err != nil {
		return nil, err
	}
	// MAPPING RECORD KE DTO
	var categoryResponses []dto.CategoryResponse
	for _, category := range categories {
		categoryResponse := dto.CategoryResponse{
			Name:        category.Name,
			Description: category.Description,
			CreatedBy:   category.CreatedBy,
			CreatedAt:   category.CreatedAt,
		}
		categoryResponses = append(categoryResponses, categoryResponse)
	}

	return categoryResponses, nil
}

func (service CategoryServiceImpl) SummaryCategory(ctx context.Context) ([]dto.SummaryCategory, error) {

	var summary []dto.SummaryCategory

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		categorySum, err := service.CategoryRepository.SummaryCategory(ctx, tx)
		if err != nil {
			return err
		}

		summary = categorySum
		return nil

	})

	if err != nil {
		return summary, errors.New("error transactional")
	}

	return summary, nil
}
