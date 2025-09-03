package config

import (
	"go-fitbyte/src/api/routes"
	"go-fitbyte/src/pkg/book"

	"gorm.io/gorm"
)

// InitServices initializes all application services
func InitServices(db *gorm.DB) routes.Services {
	// Initialize repositories
	bookRepo := book.NewRepo(db)
	// userRepo := user.NewRepo(db) // Future repo
	// authRepo := auth.NewRepo(db) // Future repo

	// Initialize services
	bookService := book.NewService(bookRepo)
	// userService := user.NewService(userRepo) // Future service
	// authService := auth.NewService(authRepo) // Future service

	return routes.Services{
		BookService: bookService,
		// UserService: userService, // Future service
		// AuthService: authService, // Future service
	}
}