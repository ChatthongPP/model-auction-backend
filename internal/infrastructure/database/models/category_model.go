package models

import (
	"time"

	"gorm.io/gorm"
)

type CategoryModel struct {
	ID        int            `gorm:"column:id;primary_key" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Image     []string       `gorm:"column:image" json:"image"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (*CategoryModel) TableName() string {
	return "categories"
}
