package dto

type ItemRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	CategoryID  uint    `json:"category_id"`
	SupplierID  uint    `json:"supplier_id"`
	//CreatedBy   uint    `json:"created_by"`
}
