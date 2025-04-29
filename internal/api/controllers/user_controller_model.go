package controllers

import "time"

type UserRequest struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	CitizenID   string `json:"citizen_id"`
	RoleID      uint   `json:"role_id"`
}

type UserResponse struct {
	ID          uint      `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"omitempty"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CitizenID   string    `json:"citizen_id"`
	RoleID      uint      `json:"role_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
