package domain

import "time"

type CategoriesData struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Image     string     `json:"image"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
