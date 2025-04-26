package controllers

import (
	"net/http"

	"backend-service/internal/domain"
	"backend-service/pkg/utilities/responses"

	"github.com/labstack/echo/v4"
)

// @Summary		Create Item selling price
// @Description	Create Item selling price for advance schedule price
// @Security		Authorization
// @Tags			Items
// @Accept			json
// @Produce		json
// @Router			/v1/items/schedule-selling-price [post]
//
// @Param			[]usecase.SellingItemDTO	body		[]usecase.SellingItemDTO	true	"Selling price item list"
// @Success		201							{object}	responses.Response
// @Failure		400							{object}	responses.ErrorResponse
// @Failure		500							{object}	responses.ErrorResponse
func (h *Controller) Hello(c echo.Context) error {
	var studentDTO StudentDTO
	if err := c.Bind(&studentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	student := domain.Student{
		FirstName: studentDTO.FirstName,
		LastName:  studentDTO.LastName,
		Email:     studentDTO.Email,
	}

	err := h.uc.CreateStudent(student)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	} else {
		return c.JSON(http.StatusCreated, responses.Ok(http.StatusCreated, "Successfully save selling items", nil))
	}
}
