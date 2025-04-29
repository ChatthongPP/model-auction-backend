package domain

type CategoryInterface interface {
	GetCategories() ([]*Category, error)
}
