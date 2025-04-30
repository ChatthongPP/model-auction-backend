package domain

type ProductInterface interface {
	CreateProduct(product *Product) error
}
