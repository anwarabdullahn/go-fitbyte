package user

import (
	"go-fitbyte/src/pkg/entities"

	"gorm.io/gorm"
)

// Repository interface allows us to access the CRUD Operations here.
type Repository interface {
	UpdateProfile(user *entities.User) (*entities.User, error)
	FetchUserById(ID uint) (*entities.User, error)
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
