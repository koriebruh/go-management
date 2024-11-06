package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"type:text"`
	CreatedBy   uint   `gorm:"not null"`
	// Relationship to Admin
	Admin Admin `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
