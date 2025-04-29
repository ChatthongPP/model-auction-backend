package usecase

import "backend-service/internal/domain"

func (uc *Usecase) GetCategories() ([]*domain.Category, error) {
	categories, err := uc.categoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
