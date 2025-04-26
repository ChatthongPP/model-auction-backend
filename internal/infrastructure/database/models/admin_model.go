package models

import "time"

type AdminModel struct {
	Id          int        `gorm:"column:id;primary_key"`
	FirstName   string     `gorm:"column:firstname"`
	LastName    string     `gorm:"column:lastname"`
	MailID      int        `gorm:"column:mail_id"`
	PasswordID  int        `gorm:"column:password_id"`
	Tel         string     `gorm:"column:tel"`
	CompanyID   int        `gorm:"column:company_id"`
	CompanyName CompanyModel `gorm:"foreignKey:company_id"`
	DeletedAt   *time.Time `gorm:"column:is_delete"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	Email        LoginModel `gorm:"foreignKey:MailID"`
	Password    LoginModel `gorm:"foreignKey:PasswordID"`
}

// TableName sets the table name for AdminModel.
func (AdminModel) TableName() string {
	return "admin_info"
}
