package user

import (
	"go-fitbyte/src/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	FetchProfile() (*entities.User, error)
	UpdateProfile(user *entities.User) (*entities.User, error)
	FetchUserById(ID uint) (*entities.User, error)
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

func (s *service) FetchProfile() (*entities.User, error) {
	return s.repository.ReadProfile()
}

// UpdateBook is a service layer that helps update books in BookShop
func (s *service) UpdateProfile(user *entities.User) (*entities.User, error) {
	return s.repository.UpdateProfile(user)
}

func (s *service) FetchUserById(ID uint) (*entities.User, error) {
	return s.repository.FetchUserById(ID)
}
