package presenter

import (
	"go-fitbyte/src/pkg/entities"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Activity is the presenter object which will be passed in the response by Handler
type Activity struct {
	ID                uint      `json:"activityId"`
	ActivityType      string    `json:"activityType"`
	DoneAt            time.Time `json:"doneAt"`
	DurationInMinutes int       `json:"durationInMinutes"`
	CaloriesBurned    int       `json:"caloriesBurned"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

// ActivitySuccessResponse is the singular SuccessResponse that will be passed in the response by Handler
func ActivitySuccessResponse(data *entities.Activity) *Activity {
	activity := Activity{
		ID:                data.ID,
		ActivityType:      data.ActivityType,
		DoneAt:            data.DoneAt,
		DurationInMinutes: data.DurationInMinutes,
		CaloriesBurned:    data.CaloriesBurned,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
	}
	return &activity
}

// ActivitiesSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func ActivitiesSuccessResponse(data *[]Activity) *[]Activity {
	return data
}

// ActivityErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ActivityErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  err.Error(),
	}
}