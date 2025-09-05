package presenter

import (
	// "go-fitbyte/src/pkg/entities"

	"go-fitbyte/src/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

// Book is the presenter object which will be passed in the response by Handler
type User struct {
	ID         uint   `json:"id" gorm:"column:id"`
	Preference string `json:"preference" gorm:"column:preference"`
	WeightUnit string `json:"weightUnit" gorm:"column:weightunit"`
	HeightUnit string `json:"heightUnit" gorm:"column:heightunit"`
	Weight     int    `json:"weight" gorm:"column:weight"`
	Height     int    `json:"height" gorm:"column:height"`
	Email      string `json:"email" gorm:"column:email"`
	Name       string `json:"name" gorm:"column:name"`
	ImageURI   string `json:"imageUri" gorm:"column:imageuri"`
}

// BookSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func ProfileSuccessResponse(data *entities.User) any {
	return map[string]any{
		"preference": data.Preference,
		"weightUnit": data.WeightUnit,
		"heightUnit": data.HeightUnit,
		"weight":     data.Weight,
		"height":     data.Height,
		"email":      data.Email,
		"name":       data.Name,
		"imageUri":   data.ImageURI,
	}
}

// BookErrorResponse is the ErrorResponse that will be passed in the response by Handler
func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
