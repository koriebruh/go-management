package domain

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name        string `gorm:"size:100;not null"`
	ContactInfo string `gorm:"size:100"`
	CreatedBy   uint   `gorm:"not null"`
	// Relationship to Admin
	Admin Admin `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
