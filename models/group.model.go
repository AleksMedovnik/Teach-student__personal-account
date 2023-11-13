package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Number uint16  `gorm:"type:integer;uniqueIndex;not null"`
	Users  []*User `json:"users" gorm:"many2many:user_group"`
}

type GroupInput struct {
	Number uint16  `json:"number" validate:"required"`
	Users  []*User `json:"users"`
}

type GroupResponse struct {
	ID     uint    `json:"id"`
	Number uint16  `json:"number"`
	Users  []*User `json:"users"`
}

func FilterGroupResord(group *Group) GroupResponse {
	return GroupResponse{
		ID:     group.ID,
		Number: group.Number,
		Users:  group.Users,
	}
}

// func (group Group) FindAll() ([]Group, error) {
// 	db, err := config.GetDB()
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		var groups []Group
// 		db.Preload("Users").Find(&groups)
// 		return groups, nil
// 	}
// }