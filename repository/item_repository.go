package repository

import (
	"context"
	"gorm.io/gorm"
	"koriebruh/management/domain"
)

type ItemRepository interface {
	Create(ctx context.Context, db *gorm.DB, item domain.Item) error
	FindAllItem(ctx context.Context, db *gorm.DB) ([]domain.Item, error)
}

type ItemRepositoryImpl struct {
}

func NewItemRepository() *ItemRepositoryImpl {
	return &ItemRepositoryImpl{}
}

func (repo ItemRepositoryImpl) Create(ctx context.Context, db *gorm.DB, item domain.Item) error {

	var existingItem domain.Item

	if err := db.WithContext(ctx).Where("name", item.Name).First(&existingItem).Error; err != nil {
		// Jika errornya bukan "record not found", kembalikan error tersebut
		if err != gorm.ErrRecordNotFound {
			return err
		}
	}

	if err := db.WithContext(ctx).Create(&item).Error; err != nil {
		return err
	}

	return nil
}

func (repo ItemRepositoryImpl) FindAllItem(ctx context.Context, db *gorm.DB) ([]domain.Item, error) {
	var items []domain.Item

	err := db.WithContext(ctx).Preload("Category").
		Preload("Supplier").
		Preload("Admin").
		Find(&items).Error

	if err != nil {
		return items, err
	}

	return items, nil

}
