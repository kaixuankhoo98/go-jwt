package routes

import (
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"
	"go-jwt/services"

	"github.com/gin-gonic/gin"
)

func SetupExerciseTypeRoutes(router *gin.Engine) {
	exerciseTypeService := services.NewExerciseTypeService(initializers.DB)

	exerciseTypeController := controllers.NewExerciseTypeController(exerciseTypeService)

	router.POST("/exerciseType", middleware.RequireAuth, exerciseTypeController.CreateExerciseType)
	router.GET("/exerciseTypes", middleware.RequireAuth, exerciseTypeController.GetExerciseTypes)
	router.PUT("/exerciseType/:exerciseTypeId", middleware.RequireAuth, exerciseTypeController.UpdateExerciseType)
	router.DELETE("/exerciseType/:exerciseTypeId", middleware.RequireAuth, exerciseTypeController.DeleteExerciseType)
}
