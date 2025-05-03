package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
	ExampleRepo   *Repo
	CategoryRepo  *Repo
	ProductRepo   *Repo
	UserRepo      *Repo
	BidRepo       *Repo
	WalletLogRepo *Repo
}

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		ExampleRepo: &Repo{
			db: db,
		},
		CategoryRepo: &Repo{
			db: db,
		},
		ProductRepo: &Repo{
			db: db,
		},
		UserRepo: &Repo{
			db: db,
		},
		BidRepo: &Repo{
			db: db,
		},
		WalletLogRepo: &Repo{
			db: db,
		},
	}
}
