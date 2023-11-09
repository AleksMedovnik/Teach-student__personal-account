package models

type Group struct {
	ID        uint16  `json:"id" gorm:"primaryKey"`
	Number    uint16  `gorm:"type:integer;not null"`
	UserRefer int     `json:"user_id"`
	Users     []*User `gorm:"many2many:user_group;"`
}

type GroupInput struct {
	Number uint16 `json:"number" validate:"required"`
}

type GroupResponse struct {
	ID        uint16 `json:"id"`
	Number    uint16 `json:"number"`
	UserRefer int    `json:"user_id"`
}

func FilterGroupResord(group *Group) GroupResponse {
	return GroupResponse{
		ID:        group.ID,
		Number:    group.Number,
		UserRefer: group.UserRefer,
	}
}
