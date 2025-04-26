package models

type StudentModel struct {
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Id        int    `gorm:"column:id;primary_key"`
}

func (e *StudentModel) TableName() string {
	return "students"
}
