package main

import (
	"go-fitbyte/src/api/routes"
	"go-fitbyte/src/config"
	"go-fitbyte/src/pkg/book"
	"go-fitbyte/src/pkg/entities"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	port := viperConfig.GetString("server.port")
	db := config.NewGorm(viperConfig)

	// Auto-migrate the Book entity
	err := db.AutoMigrate(&entities.Book{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
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

	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
