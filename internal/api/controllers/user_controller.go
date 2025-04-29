package controllers

import (
	"net/http"

	"backend-service/internal/domain"
	"backend-service/pkg/utilities/auth"

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

// @Summary		User login
// @Description	Authenticate user and return JWT token
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			credentials	body	LoginRequest	true	"Login credentials"
// @Success		200	{object}	map[string]interface{}	"Login successful"
// @Failure		400	{object}	map[string]interface{}	"Invalid input"
// @Failure		401	{object}	map[string]interface{}	"Incorrect email or password"
// @Router			/api/users/login [post]
func (c *Controller) Login(ctx echo.Context) error {
	var req LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	user, err := c.uc.Login(req.Email, req.Password)
	if err != nil {
		if err == domain.ErrEmailNotFound {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		if err == domain.ErrIncorrectPassword {
			return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Login failed"})
	}

	token, err := auth.GenerateJWT(user.ID, user.Email, user.RoleID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to generate token"})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}
