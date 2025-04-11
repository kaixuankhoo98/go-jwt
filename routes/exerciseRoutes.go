package routes

import (
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"
	"go-jwt/services"

	"github.com/gin-gonic/gin"
)

func SetupExerciseRoutes(router *gin.Engine) {
	workoutReadService := services.NewWorkoutReadService(initializers.DB)
	exerciseTypeReadService := services.NewExerciseTypeReadService(initializers.DB)
	exerciseService := services.NewExerciseService(initializers.DB, workoutReadService, exerciseTypeReadService)

	exerciseController := controllers.NewExerciseController(exerciseService)

	router.POST("/exercise", middleware.RequireAuth, exerciseController.CreateExercise)
	router.GET("/exercises", middleware.RequireAuth, exerciseController.GetExercisesByWorkoutId)
}
