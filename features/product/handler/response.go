package handler

import (
	"clean-arch/features/product"
)

type ProductResponse struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	UserID      uint   `json:"user_id" form:"user_id" gorm:"index"`
	Description string `json:"description" form:"description"`
}

func CoreToResponse(data product.Core) ProductResponse {
	return ProductResponse{
		ID:          data.ID,
		Name:        data.Name,
		UserID:      data.UserID,
		Description: data.Description,
	}
}

func CoreToResponseList(data []product.Core) []ProductResponse {
	var results []ProductResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
