package dto

type ItemRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	CategoryID  uint    `json:"category_id"`
	SupplierID  uint    `json:"supplier_id"`
	//CreatedBy   uint    `json:"created_by"`
}
