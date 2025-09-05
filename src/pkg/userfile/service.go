package userfile

import (
	"go-fitbyte/src/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	UploadUserFile(file *entities.UserFile) (*entities.UserFile, error)
	GetUserFile(ID uint) (*entities.UserFile, error)
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) UploadUserFile(file *entities.UserFile) (*entities.UserFile, error) {
	return s.repository.UploadUserFile(file)
}

func (s *service) GetUserFile(userID uint) (*entities.UserFile, error) {
	return s.repository.GetUserFile(userID)
}
