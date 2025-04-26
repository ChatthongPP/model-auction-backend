package repositories

import (
	"log"

	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"
)

func (r *Repo) CreateStudent(student domain.Student) error {
	log.Printf("CreateStudent: %+v", student)
	dbRepo := models.StudentModel{
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Email:     student.Email,
	}

	if err := r.db.Create(&dbRepo).Error; err != nil {
		return err
	}

	return nil
}
