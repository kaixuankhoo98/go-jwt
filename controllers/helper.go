package controllers

import (
	"go-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthenticatedUser(c *gin.Context) *models.User {
	userIdInterface, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil
	}

	user, ok := userIdInterface.(models.User)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return nil
	}

	return &user
}
