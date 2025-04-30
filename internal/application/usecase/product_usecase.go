package usecase

import "backend-service/internal/domain"

func (uc *Usecase) CreateProduct(product *domain.Product) error {
	err := uc.productRepo.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (uc *Usecase) GetProductByID(id int) (*domain.Product, error) {
	product, err := uc.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
