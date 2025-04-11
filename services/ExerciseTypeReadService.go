package services

import (
	"errors"
	"go-jwt/models"

	"gorm.io/gorm"
)

type exerciseTypeReadService struct {
	db *gorm.DB
}

func NewExerciseTypeReadService(database *gorm.DB) IExerciseTypeReadService {
	return &exerciseTypeReadService{db: database}
}

func (ets *exerciseTypeReadService) VerifyExerciseTypeBelongsToUser(userId uint, exerciseTypeId uint) error {
	var exerciseType models.ExerciseType
	if err := ets.db.First(&exerciseType, "id = ? AND user_id = ?", exerciseTypeId, userId).Error; err != nil {
		return errors.New("exercise type not found for user")
	}
	return nil
}
