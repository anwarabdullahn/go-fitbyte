package activity

import (
	"errors"
	"time"

	"go-fitbyte/src/api/presenter"
	"go-fitbyte/src/pkg/entities"

	"github.com/google/uuid"
)

type Filter struct {
	ActivityType      *string
	DoneAtFrom        *time.Time
	DoneAtTo          *time.Time
	CaloriesBurnedMin *int
	CaloriesBurnedMax *int
	Limit             int
	Offset            int
}

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	CreateActivity(userID uuid.UUID, req *entities.CreateActivityRequest) (*entities.Activity, error)
	FetchActivities(userID uuid.UUID, filter Filter) (*[]presenter.Activity, error)
	FetchActivityByID(activityID string, userID uuid.UUID) (*entities.Activity, error)
	UpdateActivity(activityID string, userID uuid.UUID, req *entities.CreateActivityRequest) (*entities.Activity, error)
	RemoveActivity(activityID string, userID uuid.UUID) error
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// CreateActivity is a service layer that helps insert activity
func (s *service) CreateActivity(userID uuid.UUID, req *entities.CreateActivityRequest) (*entities.Activity, error) {
	// Validate activity type
	if !isValidActivityType(req.ActivityType) {
		return nil, errors.New("invalid activity type")
	}

	// Parse doneAt time
	doneAt, err := time.Parse(time.RFC3339, req.DoneAt)
	if err != nil {
		return nil, errors.New("invalid date format, use ISO 8601 format")
	}

	// Calculate calories burned
	caloriesPerMinute, exists := entities.ActivityTypeCalories[req.ActivityType]
	if !exists {
		return nil, errors.New("invalid activity type for calorie calculation")
	}
	caloriesBurned := caloriesPerMinute * req.DurationInMinutes

	// Create activity entity
	activity := &entities.Activity{
		ID:                uuid.New(),
		ActivityType:      req.ActivityType,
		DoneAt:            doneAt,
		DurationInMinutes: req.DurationInMinutes,
		CaloriesBurned:    caloriesBurned,
		UserID:            userID,
	}

	return s.repository.CreateActivity(activity)
}

// FetchActivities is a service layer that helps fetch all activities for a user
func (s *service) FetchActivities(userID uuid.UUID, filter Filter) (*[]presenter.Activity, error) {
	return s.repository.ReadActivities(userID, filter)
}

// FetchActivityByID is a service layer that helps fetch a specific activity by ID
func (s *service) FetchActivityByID(activityID string, userID uuid.UUID) (*entities.Activity, error) {
	return s.repository.ReadActivityByID(activityID, userID)
}

// UpdateActivity is a service layer that helps update activities
func (s *service) UpdateActivity(activityID string, userID uuid.UUID, req *entities.CreateActivityRequest) (*entities.Activity, error) {
	// Validate activity type
	if !isValidActivityType(req.ActivityType) {
		return nil, errors.New("invalid activity type")
	}

	// Parse doneAt time
	doneAt, err := time.Parse(time.RFC3339, req.DoneAt)
	if err != nil {
		return nil, errors.New("invalid date format, use ISO 8601 format")
	}

	// Calculate calories burned
	caloriesPerMinute, exists := entities.ActivityTypeCalories[req.ActivityType]
	if !exists {
		return nil, errors.New("invalid activity type for calorie calculation")
	}
	caloriesBurned := caloriesPerMinute * req.DurationInMinutes

	// Get existing activity
	existingActivity, err := s.repository.ReadActivityByID(activityID, userID)
	if err != nil {
		return nil, err
	}

	// Update fields
	existingActivity.ActivityType = req.ActivityType
	existingActivity.DoneAt = doneAt
	existingActivity.DurationInMinutes = req.DurationInMinutes
	existingActivity.CaloriesBurned = caloriesBurned

	return s.repository.UpdateActivity(existingActivity)
}

// RemoveActivity is a service layer that helps remove activities
func (s *service) RemoveActivity(activityID string, userID uuid.UUID) error {
	return s.repository.DeleteActivity(activityID, userID)
}

// isValidActivityType checks if the activity type is valid
func isValidActivityType(activityType string) bool {
	validTypes := entities.ValidActivityTypes()
	for _, validType := range validTypes {
		if activityType == validType {
			return true
		}
	}
	return false
}