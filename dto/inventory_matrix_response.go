package dto

type InventoryMetrics struct {
	StockStatus struct {
		HealthyStock int `json:"healthy_stock"`
		LowStock     int `json:"low_stock"`
		OutOfStock   int `json:"out_of_stock"`
	} `json:"stock_status"`
	ValueMetrics struct {
		HighestValueCategory string  `json:"highest_value_category"`
		LowestValueCategory  string  `json:"lowest_value_category"`
		AverageItemValue     float64 `json:"average_item_value"`
	} `json:"value_metrics"`
	StockDistribution struct {
		ByCategory map[string]string `json:"by_category"`
	} `json:"stock_distribution"`
}
