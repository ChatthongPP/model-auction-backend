package domain

type ProductInterface interface {
	CreateProduct(product *Product) error
	GetProductByID(id int) (*Product, error)
	GetProducts(productFilter *FilterRequest, offet int) ([]*Product, int, error)
	UpdateProduct(product *Product) error
}
