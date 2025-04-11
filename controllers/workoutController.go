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

	c.JSON(http.StatusOK, gin.H{"workouts": workouts})
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

// func (wc *workoutController) DeleteWorkout(c *gin.Context) {} // need to cascade delete
