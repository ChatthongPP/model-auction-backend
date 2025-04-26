package models

import "time"

type CompanyModel struct {
	Id          int        `gorm:"column:id"`
	CompanyName string     `gorm:"column:company_name"`
	DeletedAt   *time.Time `gorm:"column:is_delete"`
}

func (CompanyModel) TableName() string {
	return "company" // ใส่ชื่อ table ที่ถูกต้อง
}