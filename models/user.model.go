package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	ID        *uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string         `gorm:"type:varchar(100);not null"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Role      *string        `gorm:"type:varchar(50);default:'student';not null"`
	Photo     *string        `gorm:"not null;default:'default.png'"`
	Surname   string         `gorm:"type:varchar(255)"`
	Country   string         `gorm:"type:varchar(255)"`
	City      string         `gorm:"type:varchar(255)"`
	Contacts  string         `gorm:"type:text"`
	DateBirth datatypes.Date `gorm:"default:null"`
	CreatedAt *time.Time     `gorm:"not null;default:now()"`
	UpdatedAt *time.Time     `gorm:"not null;default:now()"`
}

type SignUpInput struct {
	Name            string         `json:"name" validate:"required"`
	Email           string         `json:"email" validate:"required"`
	Password        string         `json:"password" validate:"required,min=8"`
	PasswordConfirm string         `json:"passwordConfirm" validate:"required,min=8"`
	Role            string         `json:"role" validate:"required,ne='admin'"`
	Photo           string         `json:"photo"`
	Surname         string         `json:"surname" validate:"required"`
	Country         string         `json:"country"`
	City            string         `json:"city"`
	Contacts        string         `json:"contacts"`
	DateBirth       datatypes.Date `json:"date_birth" validate:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type UserResponse struct {
	ID        uuid.UUID      `json:"id,omitempty"`
	Name      string         `json:"name,omitempty"`
	Email     string         `json:"email,omitempty"`
	Role      string         `json:"role,omitempty"`
	Photo     string         `json:"photo,omitempty"`
	DateBirth datatypes.Date `json:"date_birth"`
	Surname   string         `json:"surname,omitempty"`
	Country   string         `json:"country"`
	City      string         `json:"city"`
	Contacts  string         `json:"contacts"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:        *user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Surname:   user.Surname,
		Country:   user.Country,
		City:      user.City,
		Contacts:  user.Contacts,
		Role:      *user.Role,
		Photo:     *user.Photo,
		DateBirth: user.DateBirth,
		CreatedAt: *user.CreatedAt,
		UpdatedAt: *user.UpdatedAt,
	}
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
