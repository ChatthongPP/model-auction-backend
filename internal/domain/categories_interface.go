package domain

type CategoriesInterface interface {
	GetCategories() ([]*CategoriesData, error)
}
