package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"koriebruh/management/cnf"
	"koriebruh/management/domain"
	"koriebruh/management/dto"
	"koriebruh/management/repository"
)

type SupplierService interface {
	Create(ctx context.Context, token string, request dto.SupplierRequest) error
	FindAllSupplier(ctx context.Context) ([]dto.SuppliersResponse, error)
	SupplierSummary(ctx context.Context) ([]dto.SummarySupplier, error)
}

type SupplierServiceImpl struct {
	*gorm.DB
	repository.SupplierRepository
}

func NewSupplierService(DB *gorm.DB, supplierRepository repository.SupplierRepository) *SupplierServiceImpl {
	return &SupplierServiceImpl{DB: DB, SupplierRepository: supplierRepository}
}

func (service SupplierServiceImpl) Create(ctx context.Context, token string, request dto.SupplierRequest) error {
	var requestMapping domain.Supplier

	// TAKE
	claims := &cnf.JWTClaim{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cnf.JWT_KEY), nil
	})
	if err != nil {
		return err
	}
	id := claims.UserID

	err = service.DB.Transaction(func(tx *gorm.DB) error {

		requestMapping = domain.Supplier{
			Name:        request.Name,
			ContactInfo: request.ContactInfo,
			CreatedBy:   uint(id),
		}

		err = service.SupplierRepository.Create(ctx, tx, requestMapping)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (service SupplierServiceImpl) FindAllSupplier(ctx context.Context) ([]dto.SuppliersResponse, error) {

	var suppliers []domain.Supplier
	var supplierDto []dto.SuppliersResponse

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		allSupplier, err := service.SupplierRepository.FindAllSupplier(ctx, tx)
		if err != nil {
			return err
		}

		suppliers = allSupplier
		return nil
	})

	if err != nil {
		return supplierDto, err
	}

	for _, value := range suppliers {
		values := dto.SuppliersResponse{
			Name:        value.Name,
			ContactInfo: value.ContactInfo,
		}
		supplierDto = append(supplierDto, values)
	}

	return supplierDto, nil

}

func (service SupplierServiceImpl) SupplierSummary(ctx context.Context) ([]dto.SummarySupplier, error) {

	var summary []dto.SummarySupplier

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		supplierSummary, err := service.SupplierRepository.SupplierSummary(ctx, tx)
		if err != nil {
			return err
		}

		summary = supplierSummary
		return nil
	})

	if err != nil {
		return summary, err
	}

	return summary, nil

}
