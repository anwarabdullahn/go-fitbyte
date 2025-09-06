package main

import (
	"go-fitbyte/src/api/routes"
	"log"

	"go-fitbyte/src/config"
)

func main() {
	v := config.NewViper()
	db := config.NewGorm(v)
	app := config.NewFiber(v)

	if err := config.NewSwagger(app); err != nil {
		log.Printf("Failed to initialize Swagger: %v", err)
	}

	services := config.InitServices(db)
	routes.SetupRoutes(app, v, services)

	// Run server
	port := v.GetString("server.port")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
