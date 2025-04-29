package usecase

import (
	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/repositories"
)

type Usecase struct {
	exampleRepo    domain.ExampleInterface
	categoriesRepo domain.CategoriesInterface
}

func New(repo *repositories.Repository) *Usecase {
	return &Usecase{
		exampleRepo:    repo.ExampleRepo,
		categoriesRepo: repo.CategoriesRepo,
	}
}
