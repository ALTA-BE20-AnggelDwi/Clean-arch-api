package handler

import "clean-arch/features/product"

type ProductRequest struct {
	Name        string `json:"name" form:"name"`
	UserID      uint   `json:"user_id" form:"user_id" gorm:"index"`
	Description string `json:"description" form:"description"`
}

func RequestToCore(input ProductRequest) product.Core {
	return product.Core{
		Name:        input.Name,
		UserID:      input.UserID,
		Description: input.Description,
	}
}
