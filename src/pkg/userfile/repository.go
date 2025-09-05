package userfile

import (
	"errors"
	"go-fitbyte/src/pkg/entities"

	"gorm.io/gorm"
)

// Repository interface allows us to access the CRUD Operations here.
type Repository interface {
	UploadUserFile(file *entities.UserFile) (*entities.UserFile, error)
	GetUserFile(ID uint) (*entities.UserFile, error)
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

func (r *repository) UploadUserFile(file *entities.UserFile) (*entities.UserFile, error) {
	if err := r.DB.Save(file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func (r *repository) GetUserFile(userID uint) (*entities.UserFile, error) {
	var userFile entities.UserFile
	err := r.DB.Where("user_id = ?", userID).First(&userFile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &userFile, nil
}
