package auth

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"go-fitbyte/src/pkg/entities"
)

type Service interface {
	Register(user *entities.User) (*entities.User, error)
	Login(user *entities.User) (*entities.User, error)
}

type service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{DB: db}
}

func (s *service) Register(user *entities.User) (*entities.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.ID = uuid.New()
	user.Password = string(hashed)

	if err := s.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) Login(req *entities.User) (*entities.User, error) {
	var user entities.User

	if err := s.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	return &user, nil
}
