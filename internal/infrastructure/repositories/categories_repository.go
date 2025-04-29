package repositories

import (
	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"

	lop "github.com/samber/lo/parallel"
)

func (r *Repo) GetCategories() ([]*domain.CategoriesData, error) {
	var dbCategories []*models.CategoriesModel

	err := r.db.Find(&dbCategories).Error
	if err != nil {
		return nil, err
	}

	categoriesData := lop.Map(dbCategories, func(dbCategories *models.CategoriesModel, _ int) *domain.CategoriesData {
		return &domain.CategoriesData{
			ID:    dbCategories.ID,
			Name:  dbCategories.Name,
			Image: dbCategories.Image,
		}
	})

	return categoriesData, nil
}
