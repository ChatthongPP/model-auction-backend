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

func (uc *Usecase) GetProducts(filter *domain.FilterRequest) ([]*domain.Product, int, int, error) {
	offset := (filter.CurrrentPage - 1) * filter.Limit

	products, totalCount, err := uc.productRepo.GetProducts(filter, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := (totalCount + filter.Limit - 1) / filter.Limit

	return products, totalCount, totalPages, nil
}

func (uc *Usecase) UpdateProduct(product *domain.Product) error {
	err := uc.productRepo.UpdateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (uc *Usecase) DeleteProduct(id int) error {
	err := uc.productRepo.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
