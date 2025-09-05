package activity

import (
	"go-fitbyte/src/api/presenter"
	"go-fitbyte/src/pkg/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Repository interface allows us to access the CRUD Operations for activities
type Repository interface {
	CreateActivity(activity *entities.Activity) (*entities.Activity, error)
	ReadActivities(userID uuid.UUID, filter Filter) (*[]presenter.Activity, error)
	ReadActivityByID(activityID string, userID uuid.UUID) (*entities.Activity, error)
	UpdateActivity(activity *entities.Activity) (*entities.Activity, error)
	DeleteActivity(activityID string, userID uuid.UUID) error
}

type repository struct {
	DB *gorm.DB
}

// NewRepo is the single instance repo that is being created
func NewRepo(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

// CreateActivity is a GORM repository that helps to create activities
func (r *repository) CreateActivity(activity *entities.Activity) (*entities.Activity, error) {
	if err := r.DB.Create(activity).Error; err != nil {
		return nil, err
	}
	return activity, nil
}

// ReadActivities is a GORM repository that helps to fetch activities for a user
func (r *repository) ReadActivities(userID uuid.UUID, filter Filter) (*[]presenter.Activity, error) {
	var entityActivities []entities.Activity

	q := r.DB.Model(&entities.Activity{}).Where("user_id = ?", userID)

	if filter.ActivityType != nil {
		q = q.Where("activity_type = ?", *filter.ActivityType)
	}
	if filter.DoneAtFrom != nil {
		q = q.Where("done_at >= ?", *filter.DoneAtFrom)
	}
	if filter.DoneAtTo != nil {
		q = q.Where("done_at <= ?", *filter.DoneAtTo)
	}
	if filter.CaloriesBurnedMin != nil {
		q = q.Where("calories_burned >= ?", *filter.CaloriesBurnedMin)
	}
	if filter.CaloriesBurnedMax != nil {
		q = q.Where("calories_burned <= ?", *filter.CaloriesBurnedMax)
	}

	q = q.Offset(filter.Offset).Limit(filter.Limit)

	if err := q.Find(&entityActivities).Error; err != nil {
		return nil, err
	}
	
	// Convert entities.Activity to presenter.Activity
	var activities []presenter.Activity
	for _, activity := range entityActivities {
		activities = append(activities, presenter.Activity{
			ID:                activity.ID.String(),
			ActivityType:      activity.ActivityType,
			DoneAt:            activity.DoneAt,
			DurationInMinutes: activity.DurationInMinutes,
			CaloriesBurned:    activity.CaloriesBurned,
			CreatedAt:         activity.CreatedAt,
			UpdatedAt:         activity.UpdatedAt,
		})
	}
	
	return &activities, nil
}

// ReadActivityByID is a GORM repository that helps to fetch a specific activity by ID
func (r *repository) ReadActivityByID(activityID string, userID uuid.UUID) (*entities.Activity, error) {
	var activity entities.Activity
	if err := r.DB.Where("id = ? AND user_id = ?", activityID, userID).First(&activity).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

// UpdateActivity is a GORM repository that helps to update activities
func (r *repository) UpdateActivity(activity *entities.Activity) (*entities.Activity, error) {
	if err := r.DB.Model(activity).Updates(activity).Error; err != nil {
		return nil, err
	}
	var updatedActivity entities.Activity
	if err := r.DB.First(&updatedActivity, "id = ?", activity.ID).Error; err != nil {
		return nil, err
	}
	return &updatedActivity, nil
}

// DeleteActivity is a GORM repository that helps to delete activities
func (r *repository) DeleteActivity(activityID string, userID uuid.UUID) error {
	if err := r.DB.Where("id = ? AND user_id = ?", activityID, userID).Delete(&entities.Activity{}).Error; err != nil {
		return err
	}
	return nil
}