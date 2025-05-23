package controllers

import (
	"go-jwt/models"
	"go-jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *userController {
	return &userController{userService: userService}
}

func (uc *userController) SignUp(c *gin.Context) {
	var body struct {
		Email           string
		Password        string
		ConfirmPassword string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}

	user, err := uc.userService.CreateUser(body.Email, body.Password, body.ConfirmPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully", "userId": user.ID})
}

func (uc *userController) Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	tokenString, err := uc.userService.GenerateJWT(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func (uc *userController) Logout(c *gin.Context) {
	// Clear the JWT cookie by setting MaxAge to -1
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "logout successful",
	})
}

func (uc *userController) UpdatePassword(c *gin.Context) {
	var body struct {
		Email        string
		Password     string
		NewPassword  string
		NewPassword2 string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	user, err := uc.userService.UpdateUserPassword(body.Email, body.Password, body.NewPassword, body.NewPassword2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"userId": user.ID,
	})
}

func (uc *userController) Validate(c *gin.Context) {
	user, _ := c.Get("user")

	userModel := user.(models.User)

	c.JSON(http.StatusOK, gin.H{
		"userId": userModel.ID,
		// "username": userModel.UserName,
	})
}
