package main

import (
	"go-fitbyte/src/api/routes"
	"go-fitbyte/src/config"
	"go-fitbyte/src/pkg/book"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	port := viperConfig.GetString("server.port")
	db := config.NewGorm(viperConfig)
	if err := config.NewSwagger(app); err != nil {
		log.Printf("Failed to initialize Swagger: %v", err)
	}

	// Initialize book service with GORM
	bookRepo := book.NewRepo(db)
	bookService := book.NewService(bookRepo)

	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Hello World"))
	})

	api := app.Group("/api/v1")
	routes.BookRouter(api, bookService)

	log.Fatal(app.Listen(":" + port))
}
