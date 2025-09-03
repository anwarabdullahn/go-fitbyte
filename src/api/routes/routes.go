package routes

import (
	"go-fitbyte/src/pkg/book"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, services Services) {
	// API v1 group
	api := app.Group("/api/v1")
	
	// Register all routers
	BookRouter(api, services.BookService)
	// AuthRouter(api, services.AuthService) // Future router
	// UserRouter(api, services.UserService) // Future router
}

// Services struct holds all service dependencies
type Services struct {
	BookService book.Service
	// AuthService auth.Service // Future service
	// UserService user.Service // Future service
}
