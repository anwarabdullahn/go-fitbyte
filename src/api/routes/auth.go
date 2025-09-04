package routes

import (
	"go-fitbyte/src/api/handlers"
	"go-fitbyte/src/middleware"
	"go-fitbyte/src/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service auth.Service, jwtManager *auth.JWTManager) {
	app.Post("/register", handlers.Register(service))
	app.Post("/login", handlers.Login(service, jwtManager))

	// Protected route
	protected := app.Group("/protected", middleware.JWTProtected(jwtManager.Secret()))
	protected.Get("/me", func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		return c.JSON(fiber.Map{
			"message": "success",
			"user_id": userID,
		})
	})
}
