package routes

import (
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"
	"go-jwt/services"

	"github.com/gin-gonic/gin"
)

func SetupGroupRoutes(router *gin.Engine) {
	groupService := services.NewGroupService(initializers.DB)

	groupController := controllers.NewGroupController(groupService)

	router.POST("/group", middleware.RequireAuth, groupController.CreateGroup)
	router.GET("/groups", middleware.RequireAuth, groupController.GetGroups)
	router.PUT("/group/:groupId", middleware.RequireAuth, groupController.UpdateGroup)
	router.DELETE("/group/:groupId", middleware.RequireAuth, groupController.DeleteGroup)
}
