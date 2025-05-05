package controllers

import (
	"net/http"
	"time"

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
		BidTime:   time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := h.uc.CreateBid(bid); err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, "Failed to place bid"))
	}

	return ctx.JSON(http.StatusCreated, responses.Ok(http.StatusCreated, "Created Bid successfully", bid))
}

func (h *Controller) GetBids(ctx echo.Context) error {
	var filterReq FilterRequest
	if err := ctx.Bind(&filterReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Invalid query parameters"))
	}
	if err := validator.Validate(&filterReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	filter := &domain.FilterRequest{
		CurrrentPage: filterReq.CurrrentPage,
		Limit:        filterReq.Limit,
		OrderBy:      filterReq.OrderBy,
		Order:        filterReq.Order,
		Search:       filterReq.Search,
		UserID:       filterReq.UserID,
		ProductID:    filterReq.ProductID,
	}

	bids, totalCount, totalPages, err := h.uc.GetBids(filter)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	pagination := PaginationResponse{
		CurrrentPage: filter.CurrrentPage,
		Limit:        filter.Limit,
		TotalCount:   totalCount,
		TotalPages:   totalPages,
		Data:         bids,
	}

	return ctx.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully fetched bids", pagination))
}
