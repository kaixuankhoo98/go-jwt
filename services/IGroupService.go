package services

import "go-jwt/models"

type IGroupService interface {
	CreateGroup(userId uint, name string) (uint, error)
	GetGroupsByUserId(userId uint) ([]models.GroupResponse, error)
	UpdateGroup(userId uint, groupId uint, name string) error
	DeleteGroup(userId uint, groupId uint) error
}
