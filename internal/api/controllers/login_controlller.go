package controllers

import (
	"net/http"

	"backend-service/internal/domain"
	"backend-service/pkg/utilities/responses"
	"backend-service/pkg/utilities/validator"

	"github.com/labstack/echo/v4"
)

func (h *Controller) AdminLogin(ctx echo.Context) error {
	var loginRequest Login

	if err := ctx.Bind(&loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	// Validate adminDTO
	if err := validator.Validate(loginRequest); err != nil {
		validationErrors := validator.FormatValidationErrors(err)
		return ctx.JSON(http.StatusBadRequest, validationErrors)
	}
	admin := domain.Login{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	if err := h.uc.CheckEmailAndPassword(&admin); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Invalid email or password"))
	}

	adminResponse := domain.Login{
		Id:    admin.Id,
		Email: admin.Email,
	}

	return ctx.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully admin login", adminResponse))
}
