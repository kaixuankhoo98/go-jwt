package controllers

import (
	"go-jwt/models"
	"go-jwt/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type exerciseTypeController struct {
	exerciseTypeService services.IExerciseTypeService
}

func NewExerciseTypeController(exerciseTypeService services.IExerciseTypeService) *exerciseTypeController {
	return &exerciseTypeController{exerciseTypeService: exerciseTypeService}
}

// CreateExerciseType handles POST /exerciseType and creates a new exercise type for the user.
// It requires the exercise type name and and optional group ID in the request body.
func (ec *exerciseTypeController) CreateExerciseType(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	var body struct {
		Name    string `json:"name" binding:"required"`
		GroupId *uint  `json:"groupId"`
	}
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	exerciseTypeId, err := ec.exerciseTypeService.CreateExerciseType(user.ID, body.Name, body.GroupId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exerciseTypeId": exerciseTypeId})
}

// GetExerciseTypes handles GET /exerciseTypes and returns exercise types.
// If a groupId is provided as a query parameter, it filters exercise types by that group.
func (ec *exerciseTypeController) GetExerciseTypes(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	groupIdStr := c.Query("groupId")
	var exerciseTypes []models.ExerciseTypeResponse
	var err error
	if groupIdStr == "" {
		exerciseTypes, err = ec.exerciseTypeService.GetExerciseTypesByUserId(user.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		groupId, err := strconv.ParseUint(groupIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid groupId"})
			return
		}

		exerciseTypes, err = ec.exerciseTypeService.GetExerciseTypesByGroupId(user.ID, uint(groupId))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"exerciseTypes": exerciseTypes})
}

// UpdateExerciseType handles PUT /exerciseType/:exerciseTypeId and updates the exercise type name.
// It requires the exercise type ID in the URL and the new name in the request body.
func (ec *exerciseTypeController) UpdateExerciseType(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	var body struct {
		Name string `json:"name" binding:"required"`
	}
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	exerciseTypeIdStr := c.Param("exerciseTypeId")
	exerciseTypeId, err := strconv.ParseUint(exerciseTypeIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	err = ec.exerciseTypeService.UpdateExerciseType(user.ID, uint(exerciseTypeId), body.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// DeleteExerciseType handles DELETE /exerciseType/:exerciseTypeId and deletes the exercise type.
// It requires the exercise type ID in the URL.
func (ec *exerciseTypeController) DeleteExerciseType(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	exerciseTypeIdStr := c.Param("exerciseTypeId")
	exerciseTypeId, err := strconv.ParseUint(exerciseTypeIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	err = ec.exerciseTypeService.DeleteExerciseType(user.ID, uint(exerciseTypeId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
