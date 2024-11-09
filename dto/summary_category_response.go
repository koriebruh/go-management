package dto

type SummaryCategory struct {
	CategoryID       uint    `json:"category_id"`
	CategoryName     string  `json:"category_name"`
	ItemCount        int     `json:"item_count"`
	TotalStockValue  float64 `json:"total_stock_value"`
	AverageItemPrice float64 `json:"average_item_price"`
}
