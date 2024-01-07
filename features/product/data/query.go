package data

import (
	"clean-arch/features/product"
	"errors"

	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) product.ProductDataInterface {
	return &productQuery{
		db: db,
	}
}

func (repo *productQuery) Insert(input product.Core) error {
	productInputGorm := Product{
		Name:        input.Name,
		Description: input.Description,
	}

	tx := repo.db.Create(&productInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}

func (repo *productQuery) SelectAll() ([]product.Core, error) {
	var productsDataGorm []Product

	tx := repo.db.Find(&productsDataGorm)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var productsDataCore []product.Core
	for _, value := range productsDataGorm {
		var productCore = product.Core{
			ID:          value.ID,
			Name:        value.Name,
			UserID:      value.UserID,
			Description: value.Description,
		}
		productsDataCore = append(productsDataCore, productCore)
	}

	return productsDataCore, nil
}

func (repo *productQuery) Update(id int, input product.Core) error {
	dataGorm := CoreToModelProduct(input)

	tx := repo.db.Model(&Product{}).Where("id = ?", id).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}

	return nil
}

func (repo *productQuery) Delete(id int) error {
	tx := repo.db.Delete(&Product{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}

	return nil
}

// SelectByProductID implements product.ProductDataInterface.
func (repo *productQuery) SelectByProductID(productID int) ([]product.Core, error) {
	var productsDataGorm []Product

	tx := repo.db.Preload("User").Where("id = ?", productID).Find(&productsDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var productsDataCore []product.Core
	for _, value := range productsDataGorm {
		var productCore = product.Core{
			ID:          value.ID,
			Name:        value.Name,
			UserID:      value.UserID,
			Description: value.Description,
			// menambahkan data user
			User: product.User{
				Name:        value.User.Name,
				Email:       value.User.Email,
				Address:     value.User.Address,
				PhoneNumber: value.User.PhoneNumber,
				Role:        value.User.Role,
			},
		}
		productsDataCore = append(productsDataCore, productCore)
	}
	return productsDataCore, nil
}

// SelectByUserID implements product.ProductDataInterface.
func (repo *productQuery) SelectByUserID(userID int) ([]product.Core, error) {
	var productsDataGorm []Product

	tx := repo.db.Preload("User").Where("user_id = ?", userID).Find(&productsDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var productsDataCore []product.Core
	for _, value := range productsDataGorm {
		var productCore = product.Core{
			ID:          value.ID,
			Name:        value.Name,
			UserID:      value.UserID,
			Description: value.Description,
			// menambahkan data user
			User: product.User{
				Name:        value.User.Name,
				Email:       value.User.Email,
				Address:     value.User.Address,
				PhoneNumber: value.User.PhoneNumber,
				Role:        value.User.Role,
			},
		}
		productsDataCore = append(productsDataCore, productCore)
	}
	return productsDataCore, nil
}
