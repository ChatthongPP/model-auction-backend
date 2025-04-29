package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
	ExampleRepo    *Repo
	CategoriesRepo *Repo
}

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		ExampleRepo: &Repo{
			db: db,
		},
		CategoriesRepo: &Repo{
			db: db,
		},
	}
}
