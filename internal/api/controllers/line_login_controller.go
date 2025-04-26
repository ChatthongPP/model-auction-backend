package controllers

import (
	"net/http"

	"backend-service/pkg/utilities/responses"

	"github.com/labstack/echo/v4"
)

func (c *Controller) LoginLine(ctx echo.Context) error {
	var requestData LoginLineRequest

	if err := ctx.Bind(&requestData); err != nil {
		return ctx.JSON(http.StatusBadRequest,responses.Error(http.StatusBadRequest, err.Error()))
	}

	idToken := requestData.IDToken
	if idToken == "" {
		return ctx.JSON(http.StatusBadRequest,responses.Error(http.StatusBadRequest, "idToken not found"))
	}

	profile, err := c.uc.SaveProfile(idToken)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest,responses.Error(http.StatusBadRequest, err.Error()))
	}

	return ctx.JSON(http.StatusCreated, responses.Ok(http.StatusCreated, "Successfully saved profile", profile))
}