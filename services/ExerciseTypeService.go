package services

import (
	"errors"
	"go-jwt/models"

	"gorm.io/gorm"
)

type exerciseTypeService struct {
	db *gorm.DB
}

func NewExerciseTypeService(database *gorm.DB) IExerciseTypeService {
	return &exerciseTypeService{db: database}
}

func (service *exerciseTypeService) CreateExerciseType(userId uint, name string, groupId *uint) (uint, error) {
	if userId == 0 {
		return 0, errors.New("invalid user id")
	}

	if groupId != nil {
		var group models.Group
		if err := service.db.First(&group, "id = ? AND user_id = ?", *groupId, userId).Error; err != nil {
			return 0, errors.New("invalid groupId or userId")
		}
	}

	exerciseType := models.ExerciseType{UserID: userId, Name: name, GroupID: groupId}
	if err := service.db.Create(&exerciseType).Error; err != nil {
		return 0, errors.New("failed to create exercise type")
	}

	return exerciseType.ID, nil
}

func (service *exerciseTypeService) GetExerciseTypesByUserId(userId uint) ([]models.ExerciseTypeResponse, error) {
	if userId == 0 {
		return nil, errors.New("invalid user id")
	}

	var exerciseTypes []models.ExerciseType
	if err := service.db.Where("user_id = ?", userId).Find(&exerciseTypes).Error; err != nil {
		return nil, errors.New("failed to fetch exercise types")
	}

	exerciseTypeResponses := models.NewExerciseTypeResponseList(exerciseTypes)

	return exerciseTypeResponses, nil
}

func (service *exerciseTypeService) GetExerciseTypesByGroupId(userId uint, groupId uint) ([]models.ExerciseTypeResponse, error) {
	if userId == 0 {
		return nil, errors.New("invalid user id")
	}

	var exerciseTypes []models.ExerciseType
	if err := service.db.Where("user_id = ? AND group_id = ?", userId, groupId).Find(&exerciseTypes).Error; err != nil {
		return nil, errors.New("failed to fetch exercise types")
	}

	exerciseTypeResponses := models.NewExerciseTypeResponseList(exerciseTypes)

	return exerciseTypeResponses, nil
}

func (service *exerciseTypeService) UpdateExerciseType(userId uint, exerciseTypeId uint, name string) error {
	if userId == 0 {
		return errors.New("invalid user id")
	}

	var exerciseType models.ExerciseType
	if err := service.db.First(&exerciseType, "id = ? AND user_id = ?", exerciseTypeId, userId).Error; err != nil {
		return errors.New("invalid exercise type ID or user ID")
	}

	if exerciseType.Name == name {
		return nil
	}

	exerciseType.Name = name

	if err := service.db.Save(&exerciseType).Error; err != nil {
		return errors.New("failed to update exercise type")
	}

	return nil
}

func (service *exerciseTypeService) DeleteExerciseType(userId uint, exerciseTypeId uint) error {
	if userId == 0 {
		return errors.New("invalid user id")
	}

	var exerciseType models.ExerciseType
	if err := service.db.First(&exerciseType, "id = ? AND user_id = ?", exerciseTypeId, userId).Error; err != nil {
		return errors.New("invalid exercise type ID or user ID")
	}

	if err := service.db.Delete(&exerciseType).Error; err != nil {
		return errors.New("failed to delete exercise type")
	}

	return nil
}
