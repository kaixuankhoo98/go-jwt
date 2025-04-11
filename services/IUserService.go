package services

import "go-jwt/models"

type IUserService interface {
	CreateUser(email string, password string) (*models.User, error)
	GenerateJWT(email string, password string) (string, error)
	UpdateUserPassword(email string, password string, newPassword string, newPassword2 string) (*models.User, error)
}
