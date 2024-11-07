package dto

import "time"

type CategoryResponse struct {
	Name        string
	Description string
	CreatedBy   uint
	CreatedAt   time.Time
}
