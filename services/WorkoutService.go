package services

import (
	"errors"
	"go-jwt/initializers"
	"go-jwt/models"
	"time"

	"gorm.io/gorm"
)

type workoutService struct {
	db *gorm.DB
}

func NewWorkoutService(database *gorm.DB) IWorkoutService {
	return &workoutService{db: database}
}

func (service *workoutService) CreateWorkout(userId uint) (*models.Workout, error) {
	if userId == 0 {
		return nil, errors.New("invalid user id")
	}

	workout := models.Workout{UserID: userId, StartTime: time.Now()}
	if err := service.db.Create(&workout).Error; err != nil {
		return nil, errors.New("failed to create workout")
	}

	return &workout, nil
}

func (service *workoutService) GetWorkoutById(userId uint, workoutId uint) (*models.FullWorkoutResponse, error) {
	var workout models.Workout
	if err := initializers.DB.Preload("Exercises.Sets").Preload("Exercises.ExerciseType").First(&workout, "id = ?", workoutId).Error; err != nil {
		return nil, errors.New("invalid workout ID")
	}

	workoutResponse := models.NewFullWorkoutResponse(workout)

	return &workoutResponse, nil
}

func (service *workoutService) GetOpenWorkouts(userId uint) ([]models.WorkoutResponse, error) {
	var workouts []models.Workout
	if err := initializers.DB.Where("user_id = ? AND end_time IS NULL", userId).Find(&workouts).Error; err != nil {
		return nil, errors.New("failed to fetch open workouts")
	}

	workoutResponses := models.NewWorkoutResponseList(workouts)

	return workoutResponses, nil
}

func (service *workoutService) EndWorkout(workoutId uint, endDateTime *time.Time) (*models.Workout, error) {
	var workout models.Workout
	if err := initializers.DB.First(&workout, "id = ?", workoutId).Error; err != nil {
		return nil, errors.New("invalid workout ID")
	}
	if endDateTime != nil {
		if endDateTime.Before(workout.StartTime) {
			return nil, errors.New("end time cannot be before start time")
		}
		workout.EndTime = endDateTime
	} else {
		now := time.Now()
		workout.EndTime = &now
	}

	if err := initializers.DB.Save(&workout).Error; err != nil {
		return nil, errors.New("failed to end workout")
	}

	return &workout, nil
}

func (service *workoutService) DeleteWorkout(workoutId uint) error {
	var workout models.Workout
	if err := initializers.DB.First(&workout, "id = ?", workoutId).Error; err != nil {
		return errors.New("invalid workout ID")
	}
	if err := initializers.DB.Delete(&workout).Error; err != nil {
		return errors.New("failed to delete workout")
	}
	return nil
}
