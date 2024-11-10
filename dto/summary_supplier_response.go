package dto

type SummarySupplier struct {
	SupplierName string  `json:"supplier_name"`
	TotalItems   int64   `json:"total_items"`
	TotalValue   float64 `json:"total_value"`
}
