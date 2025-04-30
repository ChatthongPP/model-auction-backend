package usecase

import "backend-service/internal/domain"

func (uc *Usecase) CreateProduct(product *domain.Product) error {
	err := uc.productRepo.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}
