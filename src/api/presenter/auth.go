package presenter

import (
	"go-fitbyte/src/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

// Response kalau register/login sukses
func UserSuccessResponse(user *entities.User) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
		},
		"error": nil,
	}
}

// Error response
func ErrorResponse(msg string) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  msg,
	}
}
