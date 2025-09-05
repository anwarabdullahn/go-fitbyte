package routes

import (
	"go-fitbyte/src/api/handlers"
	"go-fitbyte/src/pkg/activity"

	"github.com/gofiber/fiber/v2"
)

// ActivityRouter sets up the activity routes
func ActivityRouter(app fiber.Router, activityService activity.Service, jwtMiddleware fiber.Handler) {
	// Create activity group with JWT middleware
	activityGroup := app.Group("/v1/activity")
	activityGroup.Use(jwtMiddleware)

	// Activity routes
	activityGroup.Post("/", handlers.CreateActivity(activityService))
	activityGroup.Get("/", handlers.GetActivities(activityService))
	activityGroup.Get("/:id", handlers.GetActivityByID(activityService))
	activityGroup.Put("/:id", handlers.UpdateActivity(activityService))
	activityGroup.Delete("/:id", handlers.DeleteActivity(activityService))
}