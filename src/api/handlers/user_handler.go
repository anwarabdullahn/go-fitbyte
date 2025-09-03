package handlers

import (
	"net/http"

	"go-fitbyte/src/api/presenter"
	// "go-fitbyte/src/pkg/entities"
	"go-fitbyte/src/pkg/entities"
	"go-fitbyte/src/pkg/user"

	"github.com/gofiber/fiber/v2"
)

// GetBooks is handler/controller which lists all Books from the BookShop
func GetCurrentUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := service.FetchProfile()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(data))
	}
}

func UpdateProfile(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.User
		userID := 1

		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		// 3. Cari user di DB
		user, err := service.FetchUserById(uint(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		user.Preference = entities.PreferenceType(requestBody.Preference)
		user.WeightUnit = entities.WeightUnitType(requestBody.WeightUnit)
		user.HeightUnit = entities.HeightUnitType(requestBody.HeightUnit)
		user.Weight = requestBody.Weight
		user.Height = requestBody.Height

		if requestBody.Name != "" {
			user.Name = requestBody.Name
		}
		if requestBody.ImageURI != "" {
			user.ImageURI = requestBody.ImageURI
		}

		result, err := service.UpdateProfile(user)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}
