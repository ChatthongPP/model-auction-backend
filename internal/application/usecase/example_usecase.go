package usecase

import (
	"backend-service/internal/domain"
)

func (uc *Usecase) CreateStudent(student domain.Student) error {
	err := uc.exampleRepo.CreateStudent(student)
	if err != nil {
		return err
	}
	return nil
}
