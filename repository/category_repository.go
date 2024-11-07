package repository

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"koriebruh/management/domain"
	"log"
)

// ACCEPTED PARAM , find all

type CategoryRepository interface {
	Create(ctx *fiber.Ctx, db *gorm.DB, category *domain.Category) error
	FindAllCategory(ctx *fiber.Ctx, db *gorm.DB) ([]domain.Category, error)
}

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repo CategoryRepositoryImpl) Create(ctx *fiber.Ctx, db *gorm.DB, category *domain.Category) error {
	// for check already category ??
	var existingCategory domain.Category

	if err := db.WithContext(ctx.Context()).Where("name = ?", category.Name).First(&existingCategory).Error; err == nil {
		return errors.New("category already registered")
	}

	if err := db.WithContext(ctx.Context()).Create(&category).Error; err != nil {
		log.Print("eror di db bg", err)
		return err
	}

	return nil
}

func (repo CategoryRepositoryImpl) FindAllCategory(ctx *fiber.Ctx, db *gorm.DB) ([]domain.Category, error) {
	var categories []domain.Category

	result := db.WithContext(ctx.Context()).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}
