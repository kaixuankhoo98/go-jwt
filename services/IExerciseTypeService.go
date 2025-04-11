package services

import "go-jwt/models"

type IExerciseTypeService interface {
	CreateExerciseType(userId uint, name string, groupId *uint) (uint, error)
	GetExerciseTypesByUserId(userId uint) ([]models.ExerciseTypeResponse, error)
	GetExerciseTypesByGroupId(userId uint, groupId uint) ([]models.ExerciseTypeResponse, error)
	UpdateExerciseType(userId uint, exerciseTypeId uint, name string) error
	DeleteExerciseType(userId uint, exerciseTypeId uint) error
}
