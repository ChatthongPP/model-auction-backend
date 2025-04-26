package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
	ExampleRepo *Repo
	LineRepo    *Repo
	ManageAdmin *Repo
	LoginRepo * Repo
	CompanyRepo * Repo
}

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		ExampleRepo: &Repo{
			db: db,
		},
		LineRepo: &Repo{
			db: db,
		},
		ManageAdmin: &Repo{
	        db: db,
		},
		LoginRepo: &Repo{
			db: db,
		},
		CompanyRepo: &Repo{
			db: db,
		},
	}
}
