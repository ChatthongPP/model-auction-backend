package usecase

import "backend-service/internal/domain"

func (uc *Usecase) GetCategories() ([]*domain.CategoriesData, error) {
	categories, err := uc.categoriesRepo.GetCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
