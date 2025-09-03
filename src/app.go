package main

import (
	"log"

	"go-fitbyte/src/api/routes"
	"go-fitbyte/src/config"
	"go-fitbyte/src/pkg/auth"
	"go-fitbyte/src/pkg/entities"
)

func main() {
	// Load config.yaml pake Viper
	v := config.NewViper()

	// Init DB (GORM)
	db := config.NewGorm(v)

	// Migrate tabel User
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}

	// Init Auth Service
	authService := auth.NewService(db)

	// Init Fiber
	app := config.NewFiber(v)

	// Register routes
	api := app.Group("/api")
	routes.AuthRouter(api, authService)

	// Run server
	port := v.GetString("server.port")
	if port == "" {
		port = "3000"
	}
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
