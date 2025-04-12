package routes

import (
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"
	"go-jwt/services"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userService := services.NewUserService(initializers.DB)

	userController := controllers.NewUserController(userService)

	router.POST("/signup", userController.SignUp)
	router.POST("/login", userController.Login)
	router.POST("/logout", userController.Logout)
	router.GET("/validate", middleware.RequireAuth, userController.Validate)
	router.POST("/updatePassword", middleware.RequireAuth, userController.UpdatePassword)
}
