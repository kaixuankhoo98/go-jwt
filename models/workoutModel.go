package models

import (
	"time"

	"gorm.io/gorm"
)

type Workout struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User           `gorm:"foreignKey:UserID"`
	StartTime time.Time      `gorm:"not null"`
	EndTime   *time.Time     `gorm:""`
	Exercises []Exercise     `gorm:"foreignKey:WorkoutID"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Full Workout Includes Exercises and Sets
type FullWorkoutResponse struct {
	ID        uint               `json:"id"`
	StartTime time.Time          `json:"startTime"`
	EndTime   *time.Time         `json:"endTime"`
	Exercises []ExerciseResponse `json:"exercises"`
}

func NewFullWorkoutResponse(workout Workout) FullWorkoutResponse {
	exercises := make([]ExerciseResponse, len(workout.Exercises))
	for i, ex := range workout.Exercises {
		exercises[i] = NewExerciseResponse(ex)
	}

	return FullWorkoutResponse{
		ID:        workout.ID,
		StartTime: workout.StartTime,
		EndTime:   workout.EndTime,
		Exercises: exercises,
	}
}

type WorkoutResponse struct {
	ID        uint       `json:"id"`
	StartTime time.Time  `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
}

func NewWorkoutResponse(workout Workout) WorkoutResponse {
	return WorkoutResponse{
		ID:        workout.ID,
		StartTime: workout.StartTime,
		EndTime:   workout.EndTime,
	}
}
func NewWorkoutResponseList(workouts []Workout) []WorkoutResponse {
	workoutResponses := make([]WorkoutResponse, len(workouts))
	for i, workout := range workouts {
		workoutResponses[i] = NewWorkoutResponse(workout)
	}
	return workoutResponses
}
