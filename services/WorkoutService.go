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

func (service *workoutService) GetWorkoutsByDateRange(userId uint, startDateTime time.Time, endDateTime *time.Time) ([]models.WorkoutResponse, error) {
	var workouts []models.Workout
	query := service.db.Where("user_id = ? AND start_time >= ?", userId, startDateTime)

	if endDateTime != nil {
		if startDateTime.After(*endDateTime) {
			return nil, errors.New("start time cannot be after end time")
		}
		twoYearsLater := startDateTime.AddDate(2, 0, 0)
		if endDateTime.After(twoYearsLater) {
			return nil, errors.New("cannot request more than 2 years of data")
		}
		query = query.Where("start_time <= ?", *endDateTime)
	}

	if err := query.Find(&workouts).Error; err != nil {
		return nil, errors.New("failed to fetch workouts")
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

func (service *workoutService) DeleteWorkout(userId uint, workoutId uint) error {
	var workout models.Workout
	if err := service.db.First(&workout, "id = ? AND user_id = ?", workoutId, userId).Error; err != nil {
		return errors.New("workout not found or does not belong to user")
	}
	if err := service.db.Delete(&workout).Error; err != nil {
		return errors.New("failed to delete workout")
	}
	return nil
}

func (service *workoutService) GetArchivedWorkouts(userId uint) ([]models.WorkoutResponse, error) {
	var workouts []models.Workout
	if err := service.db.Unscoped().Where("user_id = ? AND deleted_at IS NOT NULL", userId).Find(&workouts).Error; err != nil {
		return nil, errors.New("failed to fetch archived workouts")
	}

	workoutResponses := models.NewWorkoutResponseList(workouts)
	return workoutResponses, nil
}

func (service *workoutService) HardDeleteWorkout(userId uint, workoutId uint) error {
	// First verify the workout exists and belongs to the user
	var workout models.Workout
	if err := service.db.Unscoped().First(&workout, "id = ? AND user_id = ?", workoutId, userId).Error; err != nil {
		return errors.New("workout not found or does not belong to user")
	}

	// Start a transaction to ensure all deletes succeed or none do
	tx := service.db.Begin()
	if tx.Error != nil {
		return errors.New("failed to start transaction")
	}

	// Delete all sets associated with exercises in this workout
	if err := tx.Where("exercise_id IN (SELECT id FROM exercises WHERE workout_id = ?)", workoutId).Delete(&models.Set{}).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete sets")
	}

	// Delete all exercises in this workout
	if err := tx.Where("workout_id = ?", workoutId).Delete(&models.Exercise{}).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete exercises")
	}

	// Finally delete the workout itself
	if err := tx.Unscoped().Delete(&workout).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete workout")
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("failed to commit transaction")
	}

	return nil
}
