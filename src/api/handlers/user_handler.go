package handlers

import (
	"net/http"

	"go-fitbyte/src/api/presenter"
	"go-fitbyte/src/pkg/entities"
	"go-fitbyte/src/pkg/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// GetMe is handler/controller which lists current user
// @Summary      Get current user
// @Description Get user profile from JWT token
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Success      200   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /api/v1/user [get]
func GetMe(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// userIDStr, ok := c.Locals("user_id").(string)
		// if !ok {
		// 	return c.Status(fiber.StatusUnauthorized).
		// 		JSON(presenter.ErrorResponse("invalid token user_id"))
		// }

		// // convert string -> uint
		// uid64, err := strconv.ParseUint(userIDStr, 10, 32)

		// if err != nil {
		// 	return c.Status(fiber.StatusUnauthorized).
		// 		JSON(presenter.ErrorResponse("invalid user_id format"))
		// }
		data, err := service.FetchUserById(uint(1))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.ProfileSuccessResponse(data))
	}
}

// UpdateProfile is handler/controller which updates data of current user
// @Summary      Update a profile
// @Description  Update an existing profile (partial updates allowed)
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Param        user  body      entities.UpdateProfile   true  "User update request (partial fields allowed)"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /api/v1/user [put]
func UpdateProfile(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.UpdateProfile
		userID := 1

		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		// validasi
		if err := validate.Struct(requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// 3. Cari user di DB
		user, err := service.FetchUserById(uint(userID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		user.ID = uint(userID)
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
