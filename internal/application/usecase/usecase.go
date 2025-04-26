package usecase

import (
	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/repositories"
)

type Usecase struct {
	exampleRepo domain.ExampleInterface
	lineRepo domain.LineInterface
	manageAdmin domain.AdminManagementInterface
	loginRepo domain.LoginInterface
	companyRepo domain.CompanyListInterface
}

func New(repo *repositories.Repository) *Usecase {
	return &Usecase{
		exampleRepo: repo.ExampleRepo,
		lineRepo:    repo.LineRepo,
		manageAdmin: repo.ManageAdmin,
		loginRepo: repo.LoginRepo,
		companyRepo: repo.CompanyRepo,
	}
}
