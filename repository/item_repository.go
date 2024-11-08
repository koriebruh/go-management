package repository

import (
	"context"
	"gorm.io/gorm"
	"koriebruh/management/domain"
	"koriebruh/management/dto"
	"time"
)

type ItemRepository interface {
	Create(ctx context.Context, db *gorm.DB, item domain.Item) error
	FindAllItem(ctx context.Context, db *gorm.DB) ([]domain.Item, error)
	SummaryItem(ctx context.Context, db *gorm.DB) (dto.SummaryItem, error)
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

func (repo ItemRepositoryImpl) SummaryItem(ctx context.Context, db *gorm.DB) (dto.SummaryItem, error) {
	var summary dto.SummaryItem

	// Calculate TotalItems, TotalStockValue, and AverageItemPrice
	if err := db.WithContext(ctx).Model(&domain.Item{}).Select(
		"COUNT(*) AS total_items",
		"SUM(price * quantity) AS total_stock_value",
		"AVG(price) AS average_item_price",
	).Scan(&summary).Error; err != nil {
		return dto.SummaryItem{}, err
	}

	// Calculate TotalCategories
	if err := db.WithContext(ctx).Model(&domain.Item{}).
		Select("COUNT(DISTINCT category_id) AS total_categories").
		Scan(&summary.TotalCategories).Error; err != nil {
		return dto.SummaryItem{}, err
	}

	// Calculate TotalSuppliers
	if err := db.WithContext(ctx).Model(&domain.Item{}).
		Select("COUNT(DISTINCT supplier_id) AS total_suppliers").
		Scan(&summary.TotalSuppliers).Error; err != nil {
		return dto.SummaryItem{}, err
	}

	// Set UpdatedAt to the current time
	summary.UpdatedAt = time.Now()

	return summary, nil
}
