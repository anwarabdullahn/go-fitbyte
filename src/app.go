package main

import (
	"go-fitbyte/src/api/routes"
	"go-fitbyte/src/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	port := viperConfig.GetString("server.port")
	db := config.NewGorm(viperConfig)
	if err := config.NewSwagger(app); err != nil {
		log.Printf("Failed to initialize Swagger: %v", err)
	}

	// Health check endpoint
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Hello World"))
	})

	// Initialize all services
	services := config.InitServices(db)
	routes.SetupRoutes(app, services)

	log.Fatal(app.Listen(":" + port))
}
