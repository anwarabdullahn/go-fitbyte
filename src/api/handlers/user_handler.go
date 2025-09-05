package handlers

import (
	"net/http"
	"strconv"

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
		userIDStr, ok := c.Locals("user_id").(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).
				JSON(presenter.ErrorResponse("invalid token user_id"))
		}

		uid64, err := strconv.ParseUint(userIDStr, 10, 32)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				JSON(presenter.ErrorResponse("invalid user_id format"))
		}
		data, err := service.FetchUserById(uint(uid64))
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
// @Security BearerAuth
// @Param        user  body      entities.UpdateProfile   true  "User update request (partial fields allowed)"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /api/v1/user [put]
func UpdateProfile(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.UpdateProfile

		// 1. Ambil user_id dari token
		userIDStr, ok := c.Locals("user_id").(string)
		if !ok {
			return c.Status(http.StatusUnauthorized).
				JSON(presenter.ErrorResponse("invalid token user_id"))
		}

		uid64, err := strconv.ParseUint(userIDStr, 10, 32)
		if err != nil {
			return c.Status(http.StatusUnauthorized).
				JSON(presenter.ErrorResponse("invalid user_id format"))
		}

		// 2. Parse request body
		if err = c.BodyParser(&requestBody); err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(presenter.ErrorResponse("invalid request body: " + err.Error()))
		}

		// 3. Validasi request
		if err = validate.Struct(requestBody); err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(presenter.ErrorResponse(err.Error()))
		}

		// 4. Ambil user dari DB
		user, err := service.FetchUserById(uint(uid64))
		if err != nil {
			return c.Status(http.StatusNotFound).
				JSON(presenter.ErrorResponse("user not found"))
		}

		// 5. Update field user
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

		// 6. Simpan update ke DB
		updatedUser, err := service.UpdateProfile(user)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(presenter.ErrorResponse("failed to update profile: " + err.Error()))
		}

		return c.JSON(presenter.ProfileSuccessResponse(updatedUser))
	}
}
