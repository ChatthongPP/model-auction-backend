package controllers

import (
	"net/http"

	"backend-service/internal/domain"
	"backend-service/pkg/utilities/responses"

	"github.com/labstack/echo/v4"
	lop "github.com/samber/lo/parallel"
)

func (h *Controller) GetCategories(c echo.Context) error {
	categories, err := h.uc.GetCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	categoriesResponses := lop.Map(categories, func(categories *domain.CategoriesData, _ int) *CategoriesResponses {
		return &CategoriesResponses{
			ID:    categories.ID,
			Name:  categories.Name,
			Image: categories.Image,
		}
	})

	return c.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully fetched Categories", categoriesResponses))
}
