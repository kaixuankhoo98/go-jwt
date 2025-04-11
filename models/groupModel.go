package models

type Group struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Name   string `gorm:"not null;uniqueIndex:idx_user_group"` // Unique per user
	User   User   `gorm:"foreignKey:UserID"`
}

type GroupResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewGroupResponse(group Group) GroupResponse {
	return GroupResponse{
		ID:   group.ID,
		Name: group.Name,
	}
}

func NewGroupResponseList(groups []Group) []GroupResponse {
	res := make([]GroupResponse, len(groups))
	for i, g := range groups {
		res[i] = NewGroupResponse(g)
	}
	return res
}
