package controllers

import (
	"net/http"

	"backend-service/internal/domain"
	"backend-service/pkg/utilities/responses"
	"backend-service/pkg/utilities/validator"

	"github.com/labstack/echo/v4"
)

func (h *Controller) CreateBid(ctx echo.Context) error {
	var bidReq BidRequest
	if err := ctx.Bind(&bidReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Invalid request payload"))
	}

	if err := validator.Validate(&bidReq); err != nil {
		validationErrors := validator.FormatValidationErrors(err)
		return ctx.JSON(http.StatusBadRequest, validationErrors)
	}

	bid := &domain.Bid{
		ProductID: bidReq.ProductID,
		UserID:    bidReq.UserID,
		BidAmount: bidReq.BidAmount,
		BidTime:   *currentTime(),
	}

	err := h.uc.CreateBid(bid)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, "Failed to place bid"))
	}

	return ctx.JSON(http.StatusCreated, responses.Ok(http.StatusCreated, "Created Bid successfully", bid))
}
