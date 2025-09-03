package entities

import (
	"time"

	"gorm.io/gorm"
)

// User Constructs your User model under entities.
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string         `json:"email" gorm:"type:varchar(255);not null"`
	Password  string         `json:"-" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// CreateUserRequest struct is used to parse Create Requests for Users
type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// UpdateUserRequest struct is used to parse Update Requests for Users (partial updates)
type UpdateUserRequest struct {
	ID       uint    `json:"id" binding:"required"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	Password *string `json:"password,omitempty" binding:"omitempty,min=6"`
}

// LoginRequest struct is used to parse Login Requests
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
