package dto

type ItemReport struct {
	ItemName    string  `json:"item_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int
	TotalValue  float64
}
