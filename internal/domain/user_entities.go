package domain

import (
	"time"

	"backend-service/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Email       string         `json:"email"`
	Password    string         `json:"-"`
	Gender      string         `json:"gender"`
	PhoneNumber string         `json:"phone_number"`
	Address     string         `json:"address"`
	CitizenID   string         `json:"citizen_id"`
	RoleID      uint           `json:"role_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

func MapUserToDomain(m models.User) User {
	return User{
		ID:          m.ID,
		FirstName:   m.FirstName,
		LastName:    m.LastName,
		Email:       m.Email,
		Password:    m.Password,
		Gender:      m.Gender,
		PhoneNumber: m.PhoneNumber,
		Address:     m.Address,
		CitizenID:   m.CitizenID,
		RoleID:      m.RoleID,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
