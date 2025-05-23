package controllers

import (
	"go-jwt/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type workoutController struct {
	workoutService services.IWorkoutService
}

func NewWorkoutController(workoutService services.IWorkoutService) *workoutController {
	return &workoutController{workoutService: workoutService}
}

func (wc *workoutController) CreateWorkout(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	workout, err := wc.workoutService.CreateWorkout(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"workoutId": workout.ID})
}

func (wc *workoutController) GetWorkoutById(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	workoutIdStr := c.Param("workoutId")
	workoutId, err := strconv.ParseUint(workoutIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workout ID"})
		return
	}

	workout, err := wc.workoutService.GetWorkoutById(user.ID, uint(workoutId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"workout": workout})
}

func (wc *workoutController) GetOpenWorkouts(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	workouts, err := wc.workoutService.GetOpenWorkouts(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workouts)
}

func (wc *workoutController) GetWorkoutsByDateRange(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	startDateStr := c.Query("startDate")
	endDateStr := c.Query("endDate")

	if startDateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "startDate is required"})
		return
	}
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid StartDate format, expected YYYY-MM-DD"})
		return
	}

	var endDate *time.Time
	if endDateStr != "" {
		parsedEndDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid EndDate format, expected YYYY-MM-DD"})
			return
		}
		endDate = &parsedEndDate
	}

	workouts, err := wc.workoutService.GetWorkoutsByDateRange(user.ID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workouts)
}

func (wc *workoutController) EndWorkout(c *gin.Context) {
	var body struct {
		WorkoutID   uint
		EndDateTime *time.Time
	}
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}
	workout, err := wc.workoutService.EndWorkout(body.WorkoutID, body.EndDateTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"workout": workout})
}

func (wc *workoutController) DeleteWorkout(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	workoutIdStr := c.Param("workoutId")
	workoutId, err := strconv.ParseUint(workoutIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workout ID"})
		return
	}

	err = wc.workoutService.DeleteWorkout(user.ID, uint(workoutId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (wc *workoutController) GetArchivedWorkouts(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	workouts, err := wc.workoutService.GetArchivedWorkouts(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workouts)
}

func (wc *workoutController) HardDeleteWorkout(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	workoutIdStr := c.Param("workoutId")
	workoutId, err := strconv.ParseUint(workoutIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workout ID"})
		return
	}

	err = wc.workoutService.HardDeleteWorkout(user.ID, uint(workoutId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
