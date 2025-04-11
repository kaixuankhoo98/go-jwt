package controllers

import (
	"go-jwt/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type groupController struct {
	groupService services.IGroupService
}

func NewGroupController(groupService services.IGroupService) *groupController {
	return &groupController{groupService: groupService}
}

func (gc *groupController) CreateGroup(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	var body struct {
		Name string `json:"name" binding:"required"`
	}
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}
	groupId, err := gc.groupService.CreateGroup(user.ID, body.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"groupId": groupId})
}

func (gc *groupController) GetGroups(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	groups, err := gc.groupService.GetGroupsByUserId(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"groups": groups})
}

func (gc *groupController) UpdateGroup(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	var body struct {
		Name string `json:"name" binding:"required"`
	}
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	groupIdStr := c.Param("groupId")
	groupId, err := strconv.ParseUint(groupIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	err = gc.groupService.UpdateGroup(user.ID, uint(groupId), body.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (gc *groupController) DeleteGroup(c *gin.Context) {
	user := GetAuthenticatedUser(c)
	if user == nil {
		return
	}

	groupIdStr := c.Param("groupId")
	groupId, err := strconv.ParseUint(groupIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	err = gc.groupService.DeleteGroup(user.ID, uint(groupId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
