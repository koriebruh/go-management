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
		TotalStockValue      float64 `json:"total_stock_value"` // Nilai stok keseluruhan
		TotalItems           int     `json:"total_items"`       // Total jumlah barang
	} `json:"value_metrics"`
	StockDistribution struct {
		ByCategory      map[string]string `json:"by_category"`
		TotalCategories int               `json:"total_categories"` // Jumlah kategori
		TotalSuppliers  int               `json:"total_suppliers"`  // Jumlah pemasok
	} `json:"stock_distribution"`
}
