package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"koriebruh/management/domain"
	"koriebruh/management/dto"
	"log"
)

// ACCEPTED PARAM , find all

type CategoryRepository interface {
	Create(ctx context.Context, db *gorm.DB, category *domain.Category) error
	FindAllCategory(ctx context.Context, db *gorm.DB) ([]domain.Category, error)
	SummaryCategory(ctx context.Context, db *gorm.DB) ([]dto.SummaryCategory, error)
}

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repo CategoryRepositoryImpl) Create(ctx context.Context, db *gorm.DB, category *domain.Category) error {
	// for check already category ??
	var existingCategory domain.Category

	if err := db.WithContext(ctx).Where("name = ?", category.Name).First(&existingCategory).Error; err == nil {
		return errors.New("category already registered")
	}

	if err := db.WithContext(ctx).Create(&category).Error; err != nil {
		log.Print("eror di db bg", err)
		return err
	}

	return nil
}

func (repo CategoryRepositoryImpl) FindAllCategory(ctx context.Context, db *gorm.DB) ([]domain.Category, error) {
	var categories []domain.Category

	result := db.WithContext(ctx).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (repo CategoryRepositoryImpl) SummaryCategory(ctx context.Context, db *gorm.DB) ([]dto.SummaryCategory, error) {
	var summaries []dto.SummaryCategory

	// Query untuk menghitung jumlah barang, total nilai stok, dan rata-rata harga per kategori
	err := db.WithContext(ctx).Table("categories").
		Select("categories.id AS category_id, categories.name AS category_name, COUNT(items.id) AS item_count, " +
			"COALESCE(SUM(items.price * items.quantity), 0) AS total_stock_value, " +
			"COALESCE(AVG(items.price), 0) AS average_item_price").
		Joins("LEFT JOIN items ON items.category_id = categories.id").
		Group("categories.id").
		Scan(&summaries).Error

	if err != nil {
		return nil, err
	}

	return summaries, nil
}
