package handlers

import (
	"net/http"
	"strconv"
	"time"

	"go-fitbyte/src/api/presenter"
	"go-fitbyte/src/pkg/activity"
	"go-fitbyte/src/pkg/entities"
	"go-fitbyte/src/pkg/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// validator instance
var validate = validation.NewValidator()

// extractUserID extracts and converts user ID from JWT token
func extractUserID(c *fiber.Ctx) (uuid.UUID, error) {
	userIDInterface := c.Locals("user_id")
	if userIDInterface == nil {
		return uuid.Nil, errors.New("user not authenticated")
	}

	userIDStr, ok := userIDInterface.(string)
	if !ok {
		return uuid.Nil, errors.New("invalid user ID format")
	}

	parsed, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, errors.New("invalid user ID")
	}

	return parsed, nil
}

// CreateActivity is handler/controller which creates activities
// @Summary      Create a new activity
// @Description  Add a new activity to the collection
// @Tags         Activities
// @Accept       json
// @Produce      json
// @Param        activity  body      entities.CreateActivityRequest  true  "Activity object"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Failure      401   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /api/v1/activity [post]
// @Security     BearerAuth
func CreateActivity(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.CreateActivityRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}

		// Validate request using go-playground/validator
		if err := validate.Struct(&requestBody); err != nil {
			validationErrors := validation.FormatValidationErrors(err)
			c.Status(http.StatusBadRequest)
			return c.JSON(&fiber.Map{
				"status": false,
				"error":  "Validation error",
				"detail": validationErrors,
			})
		}

		// Get user ID from JWT token
		userID, err := extractUserID(c)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}

		result, err := service.CreateActivity(userID, &requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}
		return c.JSON(presenter.ActivitySuccessResponse(result))
	}
}

// GetActivities is handler/controller which lists all activities for a user
// @Summary      Get all activities
// @Description  Retrieve a list of all activities for the authenticated user
// @Tags         Activities
// @Accept       json
// @Produce      json
// @Success      200   {object}  map[string]interface{}
// @Failure      401   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /api/v1/activity [get]
// @Security     BearerAuth
func GetActivities(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get user ID from JWT token
		userID, err := extractUserID(c)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}

		// Defaults
		limit := 5
		offset := 0

		// Parse limit
		if l := c.Query("limit"); l != "" {
			if parsed, err := strconv.Atoi(l); err == nil && parsed >= 0 {
				limit = parsed
			}
		}
		// Parse offset
		if o := c.Query("offset"); o != "" {
			if parsed, err := strconv.Atoi(o); err == nil && parsed >= 0 {
				offset = parsed
			}
		}

		var activityTypePtr *string
		if at := c.Query("activityType"); at != "" {
			// validate against enum
			valid := false
			for _, vt := range entities.ValidActivityTypes() {
				if at == vt {
					valid = true
					break
				}
			}
			if valid {
				activityTypePtr = &at
			}
		}

		var doneAtFromPtr *time.Time
		if s := c.Query("doneAtFrom"); s != "" {
			if t, err := time.Parse(time.RFC3339, s); err == nil {
				doneAtFromPtr = &t
			}
		}

		var doneAtToPtr *time.Time
		if s := c.Query("doneAtTo"); s != "" {
			if t, err := time.Parse(time.RFC3339, s); err == nil {
				doneAtToPtr = &t
			}
		}

		var calMinPtr *int
		if s := c.Query("caloriesBurnedMin"); s != "" {
			if v, err := strconv.Atoi(s); err == nil {
				calMinPtr = &v
			}
		}

		var calMaxPtr *int
		if s := c.Query("caloriesBurnedMax"); s != "" {
			if v, err := strconv.Atoi(s); err == nil {
				calMaxPtr = &v
			}
		}

		filter := activity.Filter{
			ActivityType:       activityTypePtr,
			DoneAtFrom:         doneAtFromPtr,
			DoneAtTo:           doneAtToPtr,
			CaloriesBurnedMin:  calMinPtr,
			CaloriesBurnedMax:  calMaxPtr,
			Limit:              limit,
			Offset:             offset,
		}

		fetched, err := service.FetchActivities(userID, filter)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}
		return c.JSON(presenter.ActivitiesSuccessResponse(fetched))
	}
}

// GetActivityByID is handler/controller which gets a specific activity by ID
// @Summary      Get activity by ID
// @Description  Retrieve a specific activity by its ID
// @Tags         Activities
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Activity ID"
// @Success      200   {object}  map[string]interface{}
// @Failure      401   {object}  map[string]interface{}
// @Failure      404   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /api/v1/activity/{id} [get]
// @Security     BearerAuth
func GetActivityByID(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		activityID := c.Params("id")
		if activityID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ActivityErrorResponse(errors.New("activity ID is required")))
		}

		// Get user ID from JWT token
		userID, err := extractUserID(c)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}

		result, err := service.FetchActivityByID(activityID, userID)
		if err != nil {
			c.Status(http.StatusNotFound)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}
		return c.JSON(presenter.ActivitySuccessResponse(result))
	}
}

// UpdateActivity is handler/controller which updates an activity
// @Summary      Update an activity
// @Description  Update an existing activity
// @Tags         Activities
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Activity ID"
// @Param        activity  body      entities.CreateActivityRequest  true  "Activity update object"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Failure      401   {object}  map[string]interface{}
// @Failure      404   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /api/v1/activity/{id} [put]
// @Security     BearerAuth
func UpdateActivity(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		activityID := c.Params("id")
		if activityID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ActivityErrorResponse(errors.New("activity ID is required")))
		}

		var requestBody entities.CreateActivityRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}

		// Validate request using go-playground/validator
		if err := validate.Struct(&requestBody); err != nil {
			validationErrors := validation.FormatValidationErrors(err)
			c.Status(http.StatusBadRequest)
			return c.JSON(&fiber.Map{
				"status": false,
				"error":  "Validation Error",
				"details": validationErrors,
			})
		}

		// Get user ID from JWT token
		userID, err := extractUserID(c)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}

		result, err := service.UpdateActivity(activityID, userID, &requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}
		return c.JSON(presenter.ActivitySuccessResponse(result))
	}
}

// DeleteActivity is handler/controller which deletes an activity
// @Summary      Delete an activity
// @Description  Remove an activity from the collection
// @Tags         Activities
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Activity ID"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Failure      401   {object}  map[string]interface{}
// @Failure      404   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /api/v1/activity/{id} [delete]
// @Security     BearerAuth
func DeleteActivity(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		activityID := c.Params("id")
		if activityID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ActivityErrorResponse(errors.New("activity ID is required")))
		}

		// Get user ID from JWT token
		userID, err := extractUserID(c)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}

		err = service.RemoveActivity(activityID, userID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "deleted successfully",
			"error":  nil,
		})
	}
}