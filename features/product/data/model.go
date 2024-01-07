package data

import (
	"clean-arch/features/product"

	"gorm.io/gorm"
)

// struct user gorm model
type User struct {
	gorm.Model
	Name        string
	Email       string
	Address     string
	PhoneNumber string
	Role        string
}

// Struct untuk model Product
type Product struct {
	gorm.Model
	Name        string
	UserID      uint
	Description string
	User        User `gorm:"foreignKey:UserID"`
}

func CoreToModelProduct(core product.Core) Product {
	return Product{
		Name:        core.Name,
		Description: core.Description,
	}
}
