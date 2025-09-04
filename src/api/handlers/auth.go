package handlers

import (
	"net/http"

	"go-fitbyte/src/api/presenter"
	"go-fitbyte/src/pkg/auth"
	"go-fitbyte/src/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

// Register Handler
func Register(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req entities.User
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(presenter.ErrorResponse("invalid request body"))
		}

		user, err := service.Register(&req)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(presenter.ErrorResponse(err.Error()))
		}

		return c.JSON(presenter.UserSuccessResponse(user))
	}
}

// Login Handler
func Login(service auth.Service, jwtManager *auth.JWTManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req entities.User
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(presenter.ErrorResponse("invalid request body"))
		}

		user, err := service.Login(&req)
		if err != nil {
			return c.Status(http.StatusUnauthorized).
				JSON(presenter.ErrorResponse("invalid email or password"))
		}

		// generate token
		token, err := jwtManager.Generate(user.ID.String())
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(presenter.ErrorResponse("failed to generate token"))
		}

		return c.JSON(fiber.Map{
			"user":  presenter.UserSuccessResponse(user),
			"token": token,
		})
	}
}