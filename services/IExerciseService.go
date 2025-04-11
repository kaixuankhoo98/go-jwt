package services

import "go-jwt/models"

type IExerciseService interface {
	CreateExercise(userId uint, workoutId uint, exerciseTypeId uint, isTimeBased bool) (uint, error)
	GetExercisesByWorkoutId(userId uint, workoutId uint) ([]models.ExerciseResponse, error)
}
