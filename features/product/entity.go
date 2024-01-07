package product

// struct user gorm model
type User struct {
	Name        string
	Email       string
	Address     string
	PhoneNumber string
	Role        string
}

type Core struct {
	ID          uint
	Name        string
	UserID      uint
	Description string
	User        User
}

// interface untuk Data Layer
type ProductDataInterface interface {
	Insert(input Core) error
	SelectAll() ([]Core, error)
	Update(id int, input Core) error
	Delete(id int) error
	SelectByProductID(id int) ([]Core, error)
	SelectByUserID(userID int) ([]Core, error)
}

// interface untuk Service Layer
type ProductServiceInterface interface {
	Create(input Core) error
	GetAll() ([]Core, error)
	Update(id int, input Core) error
	Delete(id int) error
	SelectByProductID(id int) ([]Core, error)
	SelectByUserID(userID int) ([]Core, error)
}
