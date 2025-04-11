package controllers

import (
	"go-jwt/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type exerciseController struct {
	exerciseService services.IExerciseService
}

func NewExerciseController(exerciseService services.IExerciseService) *exerciseController {
	return &exerciseController{exerciseService: exerciseService}
}

func (ec *exerciseController) CreateExercise(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	var body struct {
		WorkoutId      uint `json:"workoutId" binding:"required"`
		ExerciseTypeId uint `json:"exerciseTypeId" binding:"required"`
		IsTimeBased    bool `json:"isTimeBased"`
	}
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	exerciseId, err := ec.exerciseService.CreateExercise(user.ID, body.WorkoutId, body.ExerciseTypeId, body.IsTimeBased)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exerciseId": exerciseId})
}

func (ec *exerciseController) GetExercisesByWorkoutId(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	workoutIdStr := c.Query("workoutId")
	workoutId, err := strconv.ParseUint(workoutIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workoutId"})
	}

	exercises, err := ec.exerciseService.GetExercisesByWorkoutId(user.ID, uint(workoutId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercises)
}
