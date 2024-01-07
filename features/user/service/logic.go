package service

import (
	"clean-arch/features/user"
	"errors"
)

type userService struct {
	userData user.UserDataInterface
}

// dependency injection
func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
	}
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core) error {
	// logic validation
	if input.Email == "" {
		return errors.New("[validation] email harus diisi")
	}
	err := service.userData.Insert(input)
	return err
}

// GetAll implements user.UserServiceInterface.
func (service *userService) GetAll() ([]user.Core, error) {
	// logic
	// memanggil func yg ada di data layer
	results, err := service.userData.SelectAll()
	return results, err
}

// Update implements user.UserServiceInterface.
func (service *userService) Update(id int, input user.Core) error {
	// validasi
	if id <= 0 {
		return errors.New("invalid id")
	}
	// validasi inputan
	// ...
	err := service.userData.Update(id, input)
	return err
}

// Delete implements user.UserServiceInterface.
func (service *userService) Delete(id int) error {
	// validasi
	if id <= 0 {
		return errors.New("invalid id")
	}

	err := service.userData.Delete(id)
	return err
}
