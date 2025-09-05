package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Activity represents a fitness activity performed by a user
type Activity struct {
	ID                uuid.UUID      `json:"activityId" gorm:"primaryKey;type:uuid"`
	ActivityType      string         `json:"activityType" gorm:"type:varchar(50);not null"`
	DoneAt            time.Time      `json:"doneAt" gorm:"not null"`
	DurationInMinutes int            `json:"durationInMinutes" gorm:"not null;min:1"`
	CaloriesBurned    int            `json:"caloriesBurned" gorm:"not null"`
	UserID            uuid.UUID      `json:"userId" gorm:"type:uuid;not null;index"`
	CreatedAt         time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}

// CreateActivityRequest represents the request payload for creating an activity
type CreateActivityRequest struct {
	ActivityType      string `json:"activityType" validate:"required,activity_type"`
	DoneAt            string `json:"doneAt" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	DurationInMinutes int    `json:"durationInMinutes" validate:"required,min=1"`
}

// ActivityTypeCalories maps activity types to their calories per minute
var ActivityTypeCalories = map[string]int{
	"Walking":    4,
	"Yoga":       4,
	"Stretching": 4,
	"Cycling":    8,
	"Swimming":   8,
	"Dancing":    8,
	"Hiking":     10,
	"Running":    10,
	"HIIT":       10,
	"JumpRope":   10,
}

// ValidActivityTypes returns a slice of valid activity types
func ValidActivityTypes() []string {
	return []string{
		"Walking", "Yoga", "Stretching", "Cycling",
		"Swimming", "Dancing", "Hiking", "Running",
		"HIIT", "JumpRope",
	}
}