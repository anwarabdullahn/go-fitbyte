package routes

import (
	"go-fitbyte/src/api/handlers"
	"go-fitbyte/src/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service auth.Service) {
    app.Post("/register", handlers.Register(service))
    app.Post("/login", handlers.Login(service))
}
