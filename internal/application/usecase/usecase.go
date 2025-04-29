package usecase

import (
	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/repositories"
)

type Usecase struct {
	exampleRepo  domain.ExampleInterface
	categoryRepo domain.CategoryInterface
	productRepo  domain.ProductInterface
	userRepo     domain.UserInterface
}

func New(repo *repositories.Repository) *Usecase {
	return &Usecase{
		exampleRepo:  repo.ExampleRepo,
		categoryRepo: repo.CategoryRepo,
		productRepo:  repo.ProductRepo,
		userRepo:     repo.UserRepo,
	}
}
