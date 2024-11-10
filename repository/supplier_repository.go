package repository

import (
	"context"
	"gorm.io/gorm"
	"koriebruh/management/domain"
	"koriebruh/management/dto"
)

type SupplierRepository interface {
	Create(ctx context.Context, db *gorm.DB, supplier domain.Supplier) error
	FindAllSupplier(ctx context.Context, db *gorm.DB) ([]domain.Supplier, error)
	SupplierSummary(ctx context.Context, db *gorm.DB) ([]dto.SummarySupplier, error)
}

type SupplierRepositoryImpl struct {
}

func NewSupplierRepository() *SupplierRepositoryImpl {
	return &SupplierRepositoryImpl{}
}

func (repo SupplierRepositoryImpl) Create(ctx context.Context, db *gorm.DB, supplier domain.Supplier) error {

	if err := db.WithContext(ctx).Create(&supplier).Error; err != nil {
		return err
	}

	return nil
}

func (repo SupplierRepositoryImpl) FindAllSupplier(ctx context.Context, db *gorm.DB) ([]domain.Supplier, error) {
	var suppliers []domain.Supplier
	err := db.WithContext(ctx).Find(&suppliers)
	if err != nil {
		return suppliers, err.Error
	}

	return suppliers, nil
}

func (repo SupplierRepositoryImpl) SupplierSummary(ctx context.Context, db *gorm.DB) ([]dto.SummarySupplier, error) {
	var summaries []dto.SummarySupplier

	err := db.WithContext(ctx).Table("items").
		Select("suppliers.name AS supplier_name, COUNT(items.id) AS total_items, SUM(items.price * items.quantity) AS total_value").
		Joins("JOIN suppliers ON items.supplier_id = suppliers.id").
		Group("items.supplier_id").
		Scan(&summaries).Error

	if err != nil {
		return nil, err
	}

	return summaries, nil

}
