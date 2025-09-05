package main

import (
	"fmt"
	"log"
	"time"

	"go-fitbyte/src/api/routes"
	"go-fitbyte/src/config"
	"go-fitbyte/src/pkg/auth"
	"go-fitbyte/src/pkg/book"
	"go-fitbyte/src/pkg/entities"
	"go-fitbyte/src/pkg/user"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title My API
// @version 1.0
// @description This is my API with JWT auth

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Load config.yaml pake Viper
	v := config.NewViper()

	// Init DB (GORM)
	db := config.NewGorm(v)

	// Auto-migrate the Book entity
	err := db.AutoMigrate(&entities.Book{}, &entities.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Init Auth Service & Repo
	authRepo := auth.NewGormRepository(db)
	authService := auth.NewService(authRepo)

	bookRepo := book.NewRepo(db)
	bookService := book.NewService(bookRepo)

	profileRepo := user.NewRepo(db)
	profileService := user.NewService(profileRepo)

	app := config.NewFiber(v)

	app.Use(cors.New())

	// Init JWT Manager (24 jam expired)
	jwtManager := auth.NewJWTManager(v.GetString("jwt.secret"), 24*time.Hour)

	// Init Fiber

	api := app.Group("/api/v1")
	routes.BookRouter(api, bookService)
	routes.AuthRouter(api, authService, jwtManager)
	routes.ProfileRouter(api, profileService, jwtManager)

	routesList := app.GetRoutes()
	log.Printf("registered routes: %d", len(routesList))
	for _, r := range routesList {
		log.Printf("%-6s -> %s", r.Method, r.Path)
		fmt.Printf("%-6s -> %s\n", r.Method, r.Path)
	}

	app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
		CacheAge: 86400,
	}))

	fmt.Print(app.GetRoutes())

	for _, r := range app.GetRoutes() {
		log.Printf("%s -> %s", r.Method, r.Path)
		fmt.Printf("%s -> %s", r.Method, r.Path)
	}

	// Run server
	port := v.GetString("server.port")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
