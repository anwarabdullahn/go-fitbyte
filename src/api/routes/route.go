package routes

import (
	"go-fitbyte/src/pkg/activity"
	"go-fitbyte/src/pkg/auth"
	"go-fitbyte/src/pkg/user"
	"go-fitbyte/src/pkg/userfile"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, v *viper.Viper, services Services) {
	// API v1 group
	api := app.Group("/api/v1")

	// Init JWT Manager (24 jam expired)
	jwtManager := auth.NewJWTManager(v.GetString("JWT_SECRET"), 24*time.Hour)

	AuthRouter(api, services.AuthService, jwtManager)
	ProfileRouter(api, services.ProfileService, jwtManager)
	UserfileRouter(api, services.ProfileService, services.UploadFileService, jwtManager)
	ActivityRouter(api, services.ActivityService, jwtManager)

	// Register all routers
	// AuthRouter(api, services.AuthService) // Future router
	// UserRouter(api, services.UserService) // Future router
}

// Services struct holds all service dependencies
type Services struct {
	AuthService       auth.Service
	ActivityService   activity.Service
	ProfileService    user.Service
	UploadFileService userfile.Service
	// AuthService auth.Service // Future service
	// UserService user.Service // Future service
}
