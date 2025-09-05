package main

import (
	"log"
	"time"

	"go-fitbyte/src/api/middleware"
	"go-fitbyte/src/api/routes"
	"go-fitbyte/src/config"
	"go-fitbyte/src/pkg/activity"
	"go-fitbyte/src/pkg/auth"
	"go-fitbyte/src/pkg/entities"
)

func main() {
	// Load config.yaml pake Viper
	v := config.NewViper()

	// Init DB (GORM)
	db := config.NewGorm(v)

	// Migrate tables
	if err := db.AutoMigrate(&entities.User{}, &entities.Activity{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}

	// Init Auth Service & Repo
	authRepo := auth.NewGormRepository(db)
	authService := auth.NewService(authRepo)

	// Init Activity Service & Repo
	activityRepo := activity.NewRepo(db)
	activityService := activity.NewService(activityRepo)

	// Init JWT Manager (24 jam expired)
	jwtManager := auth.NewJWTManager(v.GetString("jwt.secret"), 24*time.Hour)

	// Init Fiber
	app := config.NewFiber(v)

	// Register routes
	api := app.Group("/api")
	routes.AuthRouter(api, authService, jwtManager)
	routes.ActivityRouter(api, activityService, middleware.JWTProtected(jwtManager))

	// Run server
	port := v.GetString("server.port")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
