package repositories

import (
	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"

	lop "github.com/samber/lo/parallel"
)

func (r *Repo) GetCategories() ([]*domain.Category, error) {
	var dbCategories []*models.CategoryModel

	err := r.db.Find(&dbCategories).Error
	if err != nil {
		return nil, err
	}

	categoriesData := lop.Map(dbCategories, func(dbCategory *models.CategoryModel, _ int) *domain.Category {
		return &domain.Category{
			ID:        dbCategory.ID,
			Name:      dbCategory.Name,
			Image:     dbCategory.Image,
			CreatedAt: dbCategory.CreatedAt,
			UpdatedAt: dbCategory.UpdatedAt,
		}
	})

	return categoriesData, nil
}
