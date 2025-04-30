package domain

type ProductInterface interface {
	CreateProduct(product *Product) error
	GetProductByID(id int) (*Product, error)
}
