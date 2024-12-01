package dto

type SupplierRequest struct {
	Name        string `json:"name" validate:"required"`
	ContactInfo string `json:"contact_info" validate:"required"`
	CreatedBy   uint   `json:"created_by"`
}
