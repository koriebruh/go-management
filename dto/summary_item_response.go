package dto

import "time"

type SummaryItem struct {
	TotalItems       int       `json:"total_items"`
	TotalStockValue  float64   `json:"total_stock_value"`
	AverageItemPrice float64   `json:"average_item_price"`
	TotalCategories  int       `json:"total_categories"`
	TotalSuppliers   int       `json:"total_supplier"`
	UpdatedAt        time.Time `json:"updated_at"`
}
