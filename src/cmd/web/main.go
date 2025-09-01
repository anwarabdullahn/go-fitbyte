package main

import (
	"go-fitbyte/src/internal/config"
	"log"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	port := viperConfig.GetString("server.port")
	db := config.NewGorm(viperConfig)
	validate := config.NewValidator(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Validate: validate,
		Config:   viperConfig,
	})

	log.Fatal(app.Listen(":" + port))
}
