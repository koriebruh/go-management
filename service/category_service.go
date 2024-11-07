package service

import (
	"errors"
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
}

type CategoryServiceImpl struct {
	*gorm.DB
	repository.CategoryRepository
}

func NewCategoryService(DB *gorm.DB, categoryRepository repository.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{DB: DB, CategoryRepository: categoryRepository}
}

func (service CategoryServiceImpl) Create(ctx *fiber.Ctx, request dto.CategoryRequest) error {
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

		err = service.CategoryRepository.Create(ctx, tx, &category)
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
		category, err := service.CategoryRepository.FindAllCategory(ctx, tx)
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
