package service

import (
	"clean-arch/features/product"
	"errors"
)

type productService struct {
	productData product.ProductDataInterface
}

func NewProductService(productData product.ProductDataInterface) product.ProductServiceInterface {
	return &productService{
		productData: productData,
	}
}

func (service *productService) Create(input product.Core) error {
	if input.Name == "" {
		return errors.New("[validation] Name must be filled")
	}

	return service.productData.Insert(input)
}

func (service *productService) GetAll() ([]product.Core, error) {
	return service.productData.SelectAll()
}

func (service *productService) Update(id int, input product.Core) error {
	if id <= 0 {
		return errors.New("invalid ID")
	}

	return service.productData.Update(id, input)
}

func (service *productService) Delete(id int) error {
	if id <= 0 {
		return errors.New("invalid ID")
	}

	return service.productData.Delete(id)
}

// SelectByProductID implements product.ProductServiceInterface.
func (service *productService) SelectByProductID(id int) ([]product.Core, error) {
	if id <= 0 {
		return nil, errors.New("invalid ID")
	}

	return service.productData.SelectByProductID(id)
}

// SelectByUserID implements product.ProductServiceInterface.
func (service *productService) SelectByUserID(userID int) ([]product.Core, error) {
	if userID <= 0 {
		return nil, errors.New("invalid userID")
	}

	return service.productData.SelectByUserID(userID)
}
