package routes

import (
	"go-fitbyte/src/api/handlers"
	"go-fitbyte/src/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func ProfileRouter(app fiber.Router, service user.Service) {
	app.Get("/user", handlers.GetCurrentUser(service))
	// app.Post("/user", handlers.AddBook(service))
	app.Put("/user", handlers.UpdateProfile(service))
	// app.Delete("/books", handlers.RemoveBook(service))
}
