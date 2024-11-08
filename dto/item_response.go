package dto

import "time"

type ItemResponse struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       float64      `json:"price"`
	Quantity    int          `json:"quantity"`
	Category    CategoryItem `json:"category"`
	Supplier    SupplierItem `json:"supplier"`
	CreatedBy   AdminItem    `json:"created_by"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
type CategoryItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SupplierResponse DTO
type SupplierItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// AdminResponse DTO
type AdminItem struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
