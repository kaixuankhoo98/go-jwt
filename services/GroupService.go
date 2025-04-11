package services

import (
	"errors"
	"go-jwt/models"

	"gorm.io/gorm"
)

type groupService struct {
	db *gorm.DB
}

func NewGroupService(database *gorm.DB) IGroupService {
	return &groupService{db: database}
}

func (service *groupService) CreateGroup(userId uint, name string) (uint, error) {
	if userId == 0 {
		return 0, errors.New("invalid user id")
	}

	group := models.Group{Name: name, UserID: userId}
	if err := service.db.Create(&group).Error; err != nil {
		return 0, errors.New("failed to create group")
	}

	return group.ID, nil
}

func (service *groupService) GetGroupsByUserId(userId uint) ([]models.GroupResponse, error) {
	if userId == 0 {
		return nil, errors.New("invalid user id")
	}

	var groups []models.Group
	if err := service.db.Where("user_id = ?", userId).Find(&groups).Error; err != nil {
		return nil, errors.New("failed to fetch groups")
	}

	groupsResponse := models.NewGroupResponseList(groups)

	return groupsResponse, nil
}

func (service *groupService) UpdateGroup(userId uint, groupId uint, name string) error {
	if userId == 0 {
		return errors.New("invalid user id")
	}

	var group models.Group
	if err := service.db.First(&group, "id = ? AND user_id = ?", groupId, userId).Error; err != nil {
		return errors.New("invalid group ID or user ID")
	}

	if group.Name == name {
		return nil
	}

	group.Name = name
	if err := service.db.Save(&group).Error; err != nil {
		return errors.New("failed to update group")
	}

	return nil
}

func (service *groupService) DeleteGroup(userId uint, groupId uint) error {
	if userId == 0 {
		return errors.New("invalid user id")
	}

	var group models.Group
	if err := service.db.First(&group, "id = ? AND user_id = ?", groupId, userId).Error; err != nil {
		return errors.New("invalid group ID or user ID")
	}

	if err := service.db.Delete(&group).Error; err != nil {
		return errors.New("failed to delete group")
	}

	return nil
}
