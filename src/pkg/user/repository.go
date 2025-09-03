package user

import (
	"go-fitbyte/src/api/presenter"
	"go-fitbyte/src/pkg/entities"

	"gorm.io/gorm"
)

// Repository interface allows us to access the CRUD Operations here.
type Repository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	ReadUser() (*[]presenter.User, error)
	ReadProfile() (*entities.User, error)
	UpdateProfile(user *entities.User) (*entities.User, error)
	FetchUserById(ID uint) (*entities.User, error)
	DeleteUser(ID uint) error
}
type repository struct {
	DB *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewRepo(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

// CreateBook is a GORM repository that helps to create books
func (r *repository) CreateUser(user *entities.User) (*entities.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) ReadProfile() (*entities.User, error) {
	var users entities.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

// ReadBook is a GORM repository that helps to fetch books
func (r *repository) ReadUser() (*[]presenter.User, error) {
	var users []presenter.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *repository) UpdateProfile(user *entities.User) (*entities.User, error) {
	if err := r.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) FetchUserById(ID uint) (*entities.User, error) {
	var user entities.User

	// Cari user berdasarkan ID
	if err := r.DB.First(&user, ID).Error; err != nil {
		return nil, err
	}
	// Return pointer ke user
	return &user, nil
}

// DeleteBook is a GORM repository that helps to delete books
func (r *repository) DeleteUser(ID uint) error {
	if err := r.DB.Delete(&entities.User{}, ID).Error; err != nil {
		return err
	}
	return nil
}
