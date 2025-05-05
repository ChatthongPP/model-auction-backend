package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	FirstName   string         `gorm:"size:255;not null" json:"first_name"`
	LastName    string         `gorm:"size:255;not null" json:"last_name"`
	Email       string         `gorm:"size:255;unique;not null" json:"email"`
	Password    string         `gorm:"size:255;not null" json:"-"`
	Gender      string         `gorm:"size:10" json:"gender"`
	PhoneNumber string         `gorm:"size:20" json:"phone_number"`
	Address     string         `gorm:"type:text" json:"address"`
	CitizenID   string         `gorm:"size:20;unique" json:"citizen_id"`
	RoleID      uint           `gorm:"not null" json:"role_id"` // foreign key to Role
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (*User) TableName() string {
	return "users"
}
