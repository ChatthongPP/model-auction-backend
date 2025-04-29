package controllers

import (
	"net/http"

	"backend-service/internal/domain"

	"github.com/labstack/echo/v4"
)

// @Summary		Create new user
// @Description	Create a new user account
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			user	body		UserRequest	true	"User information"
// @Success		201		{object}	map[string]interface{}	"User registered successfully"
// @Failure		400		{object}	map[string]interface{}	"Invalid request body"
// @Failure		409		{object}	map[string]interface{}	"Email already exists"
// @Failure		500		{object}	map[string]interface{}	"Failed to create user"
// @Router			/api/users [post]
func (c *Controller) CreateUser(ctx echo.Context) error {
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	user := domain.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		Gender:      req.Gender,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		CitizenID:   req.CitizenID,
		RoleID:      req.RoleID,
	}

	res, err := c.uc.CreateUser(user)
	if err != nil {

		if err == domain.ErrEmailAlreadyExists {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to create user"})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User registered successfully",
		"user":    res,
	})
}
