package dto

type CategoryReportLimit struct {
	ByCategory    string
	Description   string
	TotalItem     int
	TotalQuantity int
	TotalValue    float64
	Items         []ItemReport
}

type ItemReport struct {
	Name       string
	Quantity   int
	SupplierID *uint
}
