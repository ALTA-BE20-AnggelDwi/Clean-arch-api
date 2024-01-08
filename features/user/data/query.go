package data

import (
	"clean-arch/features/user"
	"errors"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.Core) error {
	// proses mapping dari struct entities core ke model gorm
	userInputGorm := User{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		Role:        input.Role,
	}
	// simpan ke DB
	tx := repo.db.Create(&userInputGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// SelectAll implements user.UserDataInterface.
func (repo *userQuery) SelectAll() ([]user.Core, error) {
	var usersDataGorm []User
	tx := repo.db.Find(&usersDataGorm) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	var usersDataCore []user.Core
	for _, value := range usersDataGorm {
		var userCore = user.Core{
			ID:          value.ID,
			Name:        value.Name,
			Email:       value.Email,
			Password:    value.Password,
			Address:     value.Address,
			PhoneNumber: value.PhoneNumber,
			Role:        value.Role,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		}
		usersDataCore = append(usersDataCore, userCore)
	}

	return usersDataCore, nil
}

// Update implements user.UserDataInterface.
func (repo *userQuery) Update(id int, input user.Core) error {
	dataGorm := CoreToModel(input)
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(dataGorm)
	if tx.Error != nil {
		// fmt.Println("err:", tx.Error)
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// Delete implements user.UserDataInterface.
func (repo *userQuery) Delete(id int) error {
	tx := repo.db.Delete(&User{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}
	return nil
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (data *user.Core, err error) {
	var userGorm User
	tx := repo.db.Where("email = ? and password = ?", email, password).First(&userGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	result := userGorm.ModelToCore()
	return &result, nil
}
