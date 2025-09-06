package presenter

import (
	"go-fitbyte/src/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func UserFileSuccessResponse(data *entities.UserFile) *fiber.Map {
	return &fiber.Map{
		"uri": data.URI,
	}
}

// BookErrorResponse is the ErrorResponse that will be passed in the response by Handler
func UserFileErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
