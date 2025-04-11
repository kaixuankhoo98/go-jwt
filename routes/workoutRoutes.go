package routes

import (
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"
	"go-jwt/services"

	"github.com/gin-gonic/gin"
)

func SetupWorkoutRoutes(router *gin.Engine) {
	workoutService := services.NewWorkoutService(initializers.DB)

	workoutController := controllers.NewWorkoutController(workoutService)

	router.POST("/createWorkout", middleware.RequireAuth, workoutController.CreateWorkout)
	router.GET("/workout/:workoutId", middleware.RequireAuth, workoutController.GetWorkoutById)
	router.GET("/workouts/open", middleware.RequireAuth, workoutController.GetOpenWorkouts)
	router.POST("/endWorkout", middleware.RequireAuth, workoutController.EndWorkout)
}
