package models

import "time"

type LoginModel struct {
	Id          int        `gorm:"column:id;primary_key"`
	Email       string     `gorm:"column:email"`
	Password    string     `gorm:"column:password"`
	UID         string     `gorm:"column:uid"`
	DisplayName string     `gorm:"column:displayname"`
	Picture     string     `gorm:"column:picture"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	DeletedAt   *time.Time `gorm:"column:is_delete"`
}

// TableName sets the table name for LoginModel.
func (LoginModel) TableName() string {
	return "login"
}
