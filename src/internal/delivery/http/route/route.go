package route

import (
	"go-fitbyte/src/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/v1/login", c.UserController.Register)
	c.App.Post("/api/v1/register", c.UserController.Login)
}
