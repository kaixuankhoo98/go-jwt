package services

import "go-jwt/models"

type ISetService interface {
	CreateWeightedSet(userId uint, exerciseId uint, weight float64, reps int) (*models.Set, error)
	CreateDurationSet(userId uint, exerciseId uint, duration int) (*models.Set, error)
}
