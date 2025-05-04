package controllers

import (
	"net/http"
	"time"

	"backend-service/internal/domain"
	"backend-service/pkg/utilities/responses"
	"backend-service/pkg/utilities/validator"

	"github.com/labstack/echo/v4"
)

func (h *Controller) GetWalletLogs(ctx echo.Context) error {
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
		UserID:       filterReq.UserID,
	}

	walletLogs, totalCount, totalPages, err := h.uc.GetWalletLogs(filter)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	pagination := PaginationResponse{
		CurrrentPage: filter.CurrrentPage,
		Limit:        filter.Limit,
		TotalCount:   totalCount,
		TotalPages:   totalPages,
		Data:         walletLogs,
	}

	return ctx.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully fetched wallet logs", pagination))
}

func (h *Controller) CreateWalletLog(ctx echo.Context) error {
	var walletLogReq WalletLogRequest
	if err := ctx.Bind(&walletLogReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	if err := validator.Validate(&walletLogReq); err != nil {
		validationErrors := validator.FormatValidationErrors(err)
		return ctx.JSON(http.StatusBadRequest, validationErrors)
	}

	walletLog := &domain.WalletLog{
		Amount:      walletLogReq.Amount,
		UserID:      walletLogReq.UserID,
		Status:      walletLogReq.Status,
		UpdatedByID: walletLogReq.UpdatedByID,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	if err := h.uc.CreateWalletLog(walletLog); err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully created wallet log", walletLog))
}
