package controllers

import "time"

type CategoriesDTO struct {
	ID        int        `json:"id" validate:"omitempty"`
	Name      string     `json:"name" validate:"omitempty"`
	Image     string     `json:"image" validate:"omitempty"`
	CreatedAt time.Time  `json:"created_at" validate:"omitempty"`
	UpdatedAt time.Time  `json:"updated_at" validate:"omitempty"`
	DeletedAt *time.Time `json:"deleted_at" validate:"omitempty"`
}

type CategoriesResponses struct {
	ID    int    `json:"id" validate:"omitempty"`
	Name  string `json:"name" validate:"omitempty"`
	Image string `json:"image" validate:"omitempty"`
}
