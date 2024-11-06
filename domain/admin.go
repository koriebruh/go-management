package domain

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `gorm:"size:50;not null"`
	Password string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100"`

	/// Relation
	Categories []Category `gorm:"foreignKey:CreatedBy"`
	Suppliers  []Supplier `gorm:"foreignKey:CreatedBy"`
	Items      []Item     `gorm:"foreignKey:CreatedBy"`
}
