package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"koriebruh/management/cnf"
	"koriebruh/management/domain"
	"koriebruh/management/dto"
	"koriebruh/management/repository"
	"log"
)

type ItemService interface {
	Create(ctx context.Context, token string, request dto.ItemRequest) error
	FindAllItem(ctx context.Context) ([]dto.ItemResponse, error)
	SummaryItem(ctx context.Context) (dto.SummaryItem, error)
	FindByCondition(ctx context.Context, condition string, threshold int) ([]dto.ItemResponse, error)
	InventoryMetrics(ctx context.Context) (dto.InventoryMetrics, error)
	ReportItemByCategory(ctx context.Context, categoryName string) (dto.CategoryReportLimit, error)
}

type ItemServiceImpl struct {
	repository.ItemRepository
	*gorm.DB
}

func NewItemService(itemRepository repository.ItemRepository, DB *gorm.DB) *ItemServiceImpl {
	return &ItemServiceImpl{ItemRepository: itemRepository, DB: DB}
}

func (service ItemServiceImpl) Create(ctx context.Context, token string, request dto.ItemRequest) error {

	// TAKE
	claims := &cnf.JWTClaim{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cnf.JWT_KEY), nil
	})
	if err != nil {
		return err
	}
	id := claims.UserID

	// MAPPING DTO TO ENTITY
	item := domain.Item{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Quantity:    request.Quantity,
		CategoryID:  &request.CategoryID,
		SupplierID:  &request.SupplierID,
		CreatedBy:   uint(id),
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		err = service.ItemRepository.Create(ctx, tx, item)
		if err != nil {
			log.Print("error create item")
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("transaction failed")
	}

	return nil
}

func (service ItemServiceImpl) FindAllItem(ctx context.Context) ([]dto.ItemResponse, error) {

	var items []domain.Item
	var itemResponses []dto.ItemResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		item, err := service.ItemRepository.FindAllItem(ctx, tx)
		if err != nil {
			return err
		}

		items = item

		return nil
	})

	if err != nil {
		return itemResponses, errors.New("failed get record items")
	}

	//MAPPING RECORD TO DTO
	for _, item := range items {
		itemResponse := dto.ItemResponse{
			ID:          int(item.ID),
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Quantity:    item.Quantity,
			Category:    item.Category.Name,
			Supplier:    item.Supplier.Name,
			CreatedBy:   item.Admin.Username,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
		itemResponses = append(itemResponses, itemResponse)
	}

	return itemResponses, nil
}

func (service ItemServiceImpl) SummaryItem(ctx context.Context) (dto.SummaryItem, error) {

	var summary dto.SummaryItem

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		itemInfo, err := service.ItemRepository.SummaryItem(ctx, tx)
		if err != nil {
			return err
		}

		summary = itemInfo
		return nil
	})

	if err != nil {
		return summary, errors.New("error transactional item")
	}

	return summary, nil
}

func (service ItemServiceImpl) FindByCondition(ctx context.Context, condition string, threshold int) ([]dto.ItemResponse, error) {
	var items []domain.Item
	var itemResponses []dto.ItemResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		item, err := service.ItemRepository.FindByCondition(ctx, tx, condition, threshold)
		if err != nil {
			return err
		}

		items = item

		return nil
	})

	if err != nil {
		return itemResponses, errors.New("failed get record items")
	}

	//MAPPING RECORD TO DTO
	for _, item := range items {
		itemResponse := dto.ItemResponse{
			ID:          int(item.ID),
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Quantity:    item.Quantity,
			Category:    item.Category.Name,
			Supplier:    item.Supplier.Name,
			CreatedBy:   item.Admin.Username,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
		itemResponses = append(itemResponses, itemResponse)
	}

	return itemResponses, nil
}

func (service ItemServiceImpl) InventoryMetrics(ctx context.Context) (dto.InventoryMetrics, error) {
	var metric dto.InventoryMetrics

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		metricInfo, err := service.ItemRepository.InventoryMetrics(ctx, tx)
		if err != nil {
			return err
		}

		metric = metricInfo
		return nil
	})

	if err != nil {
		return metric, errors.New("error transactional item")
	}

	return metric, nil
}

func (service ItemServiceImpl) ReportItemByCategory(ctx context.Context, categoryName string) (dto.CategoryReportLimit, error) {
	var report dto.CategoryReportLimit
	var categoryReport dto.CategoryReport

	// Mulai transaksi
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		// Ambil data laporan berdasarkan kategori
		reports, err := service.ItemRepository.ReportItemByCategory(ctx, tx, categoryName)
		if err != nil {
			return err
		}
		log.Print("log bg", reports)

		categoryReport = reports // Menyimpan data ke categoryReport
		return nil
	})

	// Jika terjadi error pada transaksi
	if err != nil {
		return report, errors.New("error transactional item")
	}

	// Pemetaan data dari categoryReport ke categoryReportLimit
	report = dto.CategoryReportLimit{
		ByCategory:    categoryReport.Data.Category.Name,
		Description:   categoryReport.Data.Category.Description,
		TotalItem:     categoryReport.Data.Summary.TotalItems,
		TotalQuantity: categoryReport.Data.Summary.TotalQuantity,
		TotalValue:    categoryReport.Data.Summary.TotalValue,
		Items:         make([]dto.ItemReport, 0), // Membuat slice kosong untuk Items
	}

	// Pemetaan item-item dari categoryReport ke report.Items
	for _, item := range categoryReport.Data.Items {
		report.Items = append(report.Items, dto.ItemReport{
			Name:       item.Name,
			Quantity:   item.Quantity,
			SupplierID: item.SupplierID,
		})
	}

	// Kembalikan report yang sudah terisi
	return report, nil
}
