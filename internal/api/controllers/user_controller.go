package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"backend-service/internal/domain"
	"backend-service/pkg/utilities/auth"
	"backend-service/pkg/utilities/responses"
	"backend-service/pkg/utilities/validator"

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
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
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
			return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
		}
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusCreated, responses.Ok(http.StatusCreated, "User registered successfully", res))
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
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	user, err := c.uc.Login(req.Email, req.Password)
	if err != nil {
		if err == domain.ErrEmailNotFound {
			return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
		}
		if err == domain.ErrIncorrectPassword {
			return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
		}
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	token, err := auth.GenerateJWT(user.ID, user.Email, user.RoleID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Login successful", map[string]interface{}{
		"token": token,
	}))
}

// @Summary		Get user profile
// @Description	Retrieve profile information of the authenticated user
// @Tags			Users
// @Accept			json
// @Produce		json
// @Success		200		{object}	map[string]interface{}	"User profile retrieved successfully"
// @Failure		401		{object}	map[string]interface{}	"Unauthorized"
// @Failure		500		{object}	map[string]interface{}	"Failed to retrieve profile"
// @Router			/api/users/profile [get]
func (c *Controller) GetProfile(ctx echo.Context) error {
	// Extract userID from JWT token in context
	userIDStr, ok := ctx.Get("userID").(string)
	if !ok || userIDStr == "" {
		return ctx.JSON(http.StatusUnauthorized, responses.Error(http.StatusUnauthorized, "Unauthorized"))
	}

	// Convert userID to uint
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	user, err := c.uc.GetUserByID(uint(userID))
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return ctx.JSON(http.StatusNotFound, responses.Error(http.StatusNotFound, err.Error()))
		}
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully fetched profile", user))
}

func (h *Controller) GetUserByID(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Invalid User ID"))
	}

	user, err := h.uc.GetUserByID(uint(userID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully fetched user", user))
}

func (h *Controller) UpdateUser(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	var userReq UserRequest
	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	if err := validator.Validate(userReq); err != nil {
		validationErrors := validator.FormatValidationErrors(err)
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	user := &domain.User{
		ID:          uint(id),
		Email:       userReq.Email,
		FirstName:   userReq.FirstName,
		LastName:    userReq.LastName,
		Gender:      userReq.Gender,
		PhoneNumber: userReq.PhoneNumber,
		Address:     userReq.Address,
		CitizenID:   userReq.CitizenID,
		RoleID:      userReq.RoleID,
		UpdatedAt:   time.Now().UTC(),
	}

	if err := h.uc.UpdateUser(user); err != nil {
		if err == domain.ErrEmailAlreadyExists {
			return c.JSON(http.StatusConflict, responses.Error(http.StatusConflict, domain.ErrEmailAlreadyExists.Error()))
		}
		return c.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully updated user", user))
}
