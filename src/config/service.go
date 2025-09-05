package config

import (
	"go-fitbyte/src/api/routes"
	"go-fitbyte/src/pkg/activity"
	"go-fitbyte/src/pkg/auth"
	"go-fitbyte/src/pkg/user"
	"go-fitbyte/src/pkg/userfile"

	"gorm.io/gorm"
)

// InitServices initializes all application services
func InitServices(db *gorm.DB) routes.Services {
	// Initialize repositories
	authRepo := auth.NewGormRepository(db)
	activityRepo := activity.NewRepo(db)
	profileRepo := user.NewRepo(db)
	uploadFileRepo := userfile.NewRepo(db)

	// Initialize services
	authService := auth.NewService(authRepo)
	activityService := activity.NewService(activityRepo)
	profileService := user.NewService(profileRepo)
	uploadFileService := userfile.NewService(uploadFileRepo)

	return routes.Services{
		AuthService:       authService,
		ActivityService:   activityService,
		ProfileService:    profileService,
		UploadFileService: uploadFileService,
	}
}
