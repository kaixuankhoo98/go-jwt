package services

import (
	"go-jwt/models"
	"time"
)

type IWorkoutService interface {
	CreateWorkout(userId uint) (*models.Workout, error)
	GetWorkoutById(userId uint, workoutId uint) (*models.FullWorkoutResponse, error)
	GetOpenWorkouts(userId uint) ([]models.WorkoutResponse, error)
	EndWorkout(workoutId uint, endDateTime *time.Time) (*models.Workout, error)
	DeleteWorkout(workoutId uint) error
}
