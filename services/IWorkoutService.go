package services

import (
	"go-jwt/models"
	"time"
)

type IWorkoutService interface {
	CreateWorkout(userId uint) (*models.Workout, error)
	GetWorkoutById(userId uint, workoutId uint) (*models.FullWorkoutResponse, error)
	GetOpenWorkouts(userId uint) ([]models.WorkoutResponse, error)
	GetWorkoutsByDateRange(userId uint, startDateTime time.Time, endDateTime *time.Time) ([]models.WorkoutResponse, error)
	EndWorkout(workoutId uint, endDateTime *time.Time) (*models.Workout, error)
	DeleteWorkout(userId uint, workoutId uint) error
	GetArchivedWorkouts(userId uint) ([]models.WorkoutResponse, error)
	HardDeleteWorkout(userId uint, workoutId uint) error
}
