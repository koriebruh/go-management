package dto

type SupplierRequest struct {
	Name        string `json:"name"`
	ContactInfo string `json:"contact_info"`
	CreatedBy   uint   `json:"created_by"`
}
