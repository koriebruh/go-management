package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"koriebruh/management/domain"
	"koriebruh/management/dto"
	"time"
)

type ItemRepository interface {
	Create(ctx context.Context, db *gorm.DB, item domain.Item) error
	FindAllItem(ctx context.Context, db *gorm.DB) ([]domain.Item, error)
	SummaryItem(ctx context.Context, db *gorm.DB) (dto.SummaryItem, error)
	FindByCondition(ctx context.Context, db *gorm.DB, condition string, threshold int) ([]domain.Item, error)
	InventoryMetrics(ctx context.Context, db *gorm.DB) (dto.InventoryMetrics, error)
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

func (repo ItemRepositoryImpl) FindByCondition(ctx context.Context, db *gorm.DB, condition string, threshold int) ([]domain.Item, error) {
	var items []domain.Item

	switch condition {
	case "under":
		if err := db.WithContext(ctx).Where("quantity < ?", threshold).Find(&items).Error; err != nil {
			return nil, err
		}
	case "over":
		if err := db.WithContext(ctx).Where("quantity > ?", threshold).Find(&items).Error; err != nil {
			return nil, err
		}
	case "equal":
		if err := db.WithContext(ctx).Where("quantity = ?", threshold).Find(&items).Error; err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid condition: must be 'under', 'over', or 'equal'")
	}

	return items, nil
}

func (repo ItemRepositoryImpl) InventoryMetrics(ctx context.Context, db *gorm.DB) (dto.InventoryMetrics, error) {
	var metrics dto.InventoryMetrics

	// 1. Stock Status
	// Hitung stok yang sehat, stok rendah, dan stok habis
	err := db.WithContext(ctx).Table("items").Select(`
        COUNT(CASE WHEN quantity >= 50 THEN 1 END) AS healthy_stock,
        COUNT(CASE WHEN quantity < 50 AND quantity > 0 THEN 1 END) AS low_stock,
        COUNT(CASE WHEN quantity = 0 THEN 1 END) AS out_of_stock
    `).Scan(&metrics.StockStatus).Error
	if err != nil {
		return metrics, err
	}

	// 2. Value Metrics
	// Hitung kategori dengan nilai stok tertinggi dan terendah, serta nilai rata-rata item
	var highestValue, lowestValue struct {
		CategoryName string
		TotalValue   float64
	}

	// Query untuk highest value category dan total stock value
	err = db.WithContext(ctx).Table("categories").
		Select("categories.name AS category_name, COALESCE(SUM(items.price * items.quantity), 0) AS total_value").
		Joins("LEFT JOIN items ON items.category_id = categories.id").
		Group("categories.id").
		Order("total_value DESC").
		Limit(1).
		Scan(&highestValue).Error
	if err != nil {
		return metrics, err
	}

	// Query untuk lowest value category
	err = db.WithContext(ctx).Table("categories").
		Select("categories.name AS category_name, COALESCE(SUM(items.price * items.quantity), 0) AS total_value").
		Joins("LEFT JOIN items ON items.category_id = categories.id").
		Group("categories.id").
		Order("total_value ASC").
		Limit(1).
		Scan(&lowestValue).Error
	if err != nil {
		return metrics, err
	}

	metrics.ValueMetrics.HighestValueCategory = highestValue.CategoryName
	metrics.ValueMetrics.LowestValueCategory = lowestValue.CategoryName

	// Hitung total stock value
	err = db.WithContext(ctx).Table("items").
		Select("COALESCE(SUM(price * quantity), 0) AS total_stock_value").
		Scan(&metrics.ValueMetrics.TotalStockValue).Error
	if err != nil {
		return metrics, err
	}

	// Hitung total items
	err = db.WithContext(ctx).Table("items").
		Select("COALESCE(SUM(quantity), 0) AS total_items").
		Scan(&metrics.ValueMetrics.TotalItems).Error
	if err != nil {
		return metrics, err
	}

	// Hitung nilai rata-rata item
	err = db.WithContext(ctx).Table("items").
		Select("COALESCE(AVG(price), 0) AS average_item_value").
		Scan(&metrics.ValueMetrics.AverageItemValue).Error
	if err != nil {
		return metrics, err
	}

	// 3. Stock Distribution
	// Hitung distribusi stok berdasarkan kategori
	var categoryDistribution []struct {
		CategoryName string
		TotalItems   int
	}

	err = db.WithContext(ctx).Table("categories").
		Select("categories.name AS category_name, COALESCE(SUM(items.quantity), 0) AS total_items").
		Joins("LEFT JOIN items ON items.category_id = categories.id").
		Group("categories.id").
		Scan(&categoryDistribution).Error
	if err != nil {
		return metrics, err
	}

	// Hitung total categories
	err = db.WithContext(ctx).Table("categories").
		Select("COUNT(id) as total_categories").
		Scan(&metrics.StockDistribution.TotalCategories).Error
	if err != nil {
		return metrics, err
	}

	// Hitung total suppliers
	err = db.WithContext(ctx).Table("suppliers").
		Select("COUNT(id) as total_suppliers").
		Scan(&metrics.StockDistribution.TotalSuppliers).Error
	if err != nil {
		return metrics, err
	}

	// Buat distribusi per kategori dalam persen
	metrics.StockDistribution.ByCategory = make(map[string]string)
	totalItems := metrics.ValueMetrics.TotalItems
	for _, category := range categoryDistribution {
		if totalItems > 0 {
			percentage := float64(category.TotalItems) / float64(totalItems) * 100
			metrics.StockDistribution.ByCategory[category.CategoryName] = fmt.Sprintf("%.2f%%", percentage)
		} else {
			metrics.StockDistribution.ByCategory[category.CategoryName] = "0%"
		}
	}

	return metrics, nil
}
