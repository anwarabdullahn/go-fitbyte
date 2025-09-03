package user

import (
	"go-fitbyte/src/api/presenter"
	"go-fitbyte/src/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertUser(user *entities.User) (*entities.User, error)
	FetchUsers() (*[]presenter.User, error)
	FetchProfile() (*entities.User, error)
	UpdateProfile(user *entities.User) (*entities.User, error)
	FetchUserById(ID uint) (*entities.User, error)
	RemoveUser(ID uint) error
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

// InsertBook is a service layer that helps insert book in BookShop
func (s *service) InsertUser(user *entities.User) (*entities.User, error) {
	return s.repository.CreateUser(user)
}

func (s *service) FetchProfile() (*entities.User, error) {
	return s.repository.ReadProfile()
}

// FetchBooks is a service layer that helps fetch all books in BookShop
func (s *service) FetchUsers() (*[]presenter.User, error) {
	return s.repository.ReadUser()
}

// UpdateBook is a service layer that helps update books in BookShop
func (s *service) UpdateProfile(user *entities.User) (*entities.User, error) {
	return s.repository.UpdateProfile(user)
}

func (s *service) FetchUserById(ID uint) (*entities.User, error) {
	return s.repository.FetchUserById(ID)
}

// RemoveBook is a service layer that helps remove books from BookShop
func (s *service) RemoveUser(ID uint) error {
	return s.repository.DeleteUser(ID)
}
