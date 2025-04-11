package controllers

import (
	"go-jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type setController struct {
	setService services.ISetService
}

func NewSetController(setService services.ISetService) *setController {
	return &setController{setService: setService}
}

// CreateWeightedSet creates a new weighted set for the authenticated user.
// It expects a JSON body with the exerciseId, weight, and reps fields.
func (sc *setController) CreateWeightedSet(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	var body struct {
		ExerciseId uint    `json:"exerciseId" binding:"required"`
		Weight     float64 `json:"weight"`
		Reps       int     `json:"reps"`
	}
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	set, err := sc.setService.CreateWeightedSet(user.ID, body.ExerciseId, body.Weight, body.Reps)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"setId": set.ID})
}

// CreateDurationSet creates a new duration set for the authenticated user.
// It expects a JSON body with the exerciseId and duration (in seconds) fields.
func (sc *setController) CreateDurationSet(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	var body struct {
		ExerciseId uint `json:"exerciseId" binding:"required"`
		Duration   int  `json:"duration" binding:"required"`
	}
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	set, err := sc.setService.CreateDurationSet(user.ID, body.ExerciseId, body.Duration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"setId": set.ID})
}
