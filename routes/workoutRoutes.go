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
	router.DELETE("/workout/:workoutId", middleware.RequireAuth, workoutController.DeleteWorkout)
	router.DELETE("/workout/:workoutId/permanent", middleware.RequireAuth, workoutController.HardDeleteWorkout)
	router.GET("/workouts/open", middleware.RequireAuth, workoutController.GetOpenWorkouts)
	router.GET("/workouts/archived", middleware.RequireAuth, workoutController.GetArchivedWorkouts)
	router.GET("/workouts", middleware.RequireAuth, workoutController.GetWorkoutsByDateRange)
	router.POST("/endWorkout", middleware.RequireAuth, workoutController.EndWorkout)
}
