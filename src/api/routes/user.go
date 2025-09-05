package routes

import (
	"go-fitbyte/src/api/handlers"
	"go-fitbyte/src/api/middleware"
	"go-fitbyte/src/pkg/auth"
	"go-fitbyte/src/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func ProfileRouter(app fiber.Router, service user.Service, jm *auth.JWTManager) {
	profile := app.Group("/user", middleware.JWTProtected(jm))
	profile.Get("", handlers.GetMe(service))
	profile.Put("", handlers.UpdateProfile(service))
}
