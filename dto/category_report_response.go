package dto

import (
	"koriebruh/management/domain"
)

type CategoryReport struct {
	Status string `json:"status"`
	Data   struct {
		Category domain.Category `json:"category"`
		Items    []domain.Item   `json:"items"`
		Summary  struct {
			TotalItems    int     `json:"total_items"`
			TotalQuantity int     `json:"total_quantity"`
			TotalValue    float64 `json:"total_value"`
		} `json:"summary"`
	} `json:"data"`
}
