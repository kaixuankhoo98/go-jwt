package routes

import (
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"
	"go-jwt/services"

	"github.com/gin-gonic/gin"
)

func SetupSetRoutes(router *gin.Engine) {
	exerciseReadService := services.NewExerciseReadService(initializers.DB)
	workoutReadService := services.NewWorkoutReadService(initializers.DB)
	setService := services.NewSetService(initializers.DB, exerciseReadService, workoutReadService)

	setController := controllers.NewSetController(setService)

	router.POST("/set/weighted", middleware.RequireAuth, setController.CreateWeightedSet)
	router.POST("/set/duration", middleware.RequireAuth, setController.CreateDurationSet)
}
