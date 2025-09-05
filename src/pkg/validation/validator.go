package validation

import (
	"go-fitbyte/src/pkg/entities"

	"github.com/go-playground/validator/v10"
)

// NewValidator creates a new validator instance with custom validations
func NewValidator() *validator.Validate {
	validate := validator.New()
	
	// Register custom activity_type validator
	validate.RegisterValidation("activity_type", validateActivityType)
	
	return validate
}

// validateActivityType validates if the activity type is one of the allowed values
func validateActivityType(fl validator.FieldLevel) bool {
	activityType := fl.Field().String()
	validTypes := entities.ValidActivityTypes()
	
	for _, validType := range validTypes {
		if activityType == validType {
			return true
		}
	}
	return false
}

// ValidationError represents a validation error with field and message
type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

// FormatValidationErrors formats validator errors into a more readable format
func FormatValidationErrors(err error) []ValidationError {
	var validationErrors []ValidationError
	
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErr {
			validationErrors = append(validationErrors, ValidationError{
				Field:   e.Field(),
				Tag:     e.Tag(),
				Value:   e.Param(),
				Message: getValidationMessage(e),
			})
		}
	}
	
	return validationErrors
}

// getValidationMessage returns a user-friendly validation message
func getValidationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "min":
		return e.Field() + " must be at least " + e.Param()
	case "activity_type":
		return e.Field() + " must be one of: Walking, Yoga, Stretching, Cycling, Swimming, Dancing, Hiking, Running, HIIT, JumpRope"
	case "datetime":
		return e.Field() + " must be a valid ISO 8601 datetime format (e.g., 2024-01-15T10:30:00Z)"
	default:
		return e.Field() + " is invalid"
	}
}