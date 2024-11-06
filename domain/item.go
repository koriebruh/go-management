package domain

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string  `gorm:"size:100;not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	Quantity    int     `gorm:"not null"`
	CategoryID  *uint   `gorm:""`
	SupplierID  *uint   `gorm:""`
	CreatedBy   uint    `gorm:"not null"`
	// Relationships to other tables
	Category Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Supplier Supplier `gorm:"foreignKey:SupplierID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Admin    Admin    `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
