package services

import (
	"errors"
	"go-jwt/models"

	"gorm.io/gorm"
)

type workoutReadService struct {
	db *gorm.DB
}

func NewWorkoutReadService(database *gorm.DB) IWorkoutReadService {
	return &workoutReadService{db: database}
}

func (wrs *workoutReadService) VerifyWorkoutBelongsToUser(userId uint, workoutId uint) error {
	var workout models.Workout
	if err := wrs.db.First(&workout, "id = ? AND user_id = ?", workoutId, userId).Error; err != nil {
		return errors.New("workout not found for user")
	}
	return nil
}
