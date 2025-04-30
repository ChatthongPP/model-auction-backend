package controllers

import (
	"net/http"
	"strconv"

	"backend-service/internal/domain"
	"backend-service/pkg/utilities/responses"
	"backend-service/pkg/utilities/validator"

	"github.com/labstack/echo/v4"
)

func (h *Controller) CreateProduct(ctx echo.Context) error {
	var productReq ProductRequest
	if err := ctx.Bind(&productReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	if err := validator.Validate(&productReq); err != nil {
		validationErrors := validator.FormatValidationErrors(err)
		return ctx.JSON(http.StatusBadRequest, validationErrors)
	}
	product := &domain.Product{
		Name:                productReq.Name,
		Description:         productReq.Description,
		CategoryID:          productReq.CategoryID,
		SellerID:            productReq.SellerID,
		ActualPrice:         productReq.ActualPrice,
		StartingBidPrice:    productReq.StartingBidPrice,
		CurrentBidPrice:     productReq.StartingBidPrice, // เริ่มต้น CurrentBidPrice เท่ากับ StartingBidPrice
		MinimumBidIncrement: productReq.MinimumBidIncrement,
		ShippingPrice:       productReq.ShippingPrice,
		ServiceFee:          productReq.ServiceFee, //
		AuctionStartTime:    productReq.AuctionStartTime,
		AuctionEndTime:      productReq.AuctionEndTime,
		Status:              productReq.Status,
	}

	if err := h.uc.CreateProduct(product); err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully created product", product))
}

func (h *Controller) GetProductByID(ctx echo.Context) error {
	productID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Invalid Product ID"))
	}

	product, err := h.uc.GetProductByID(productID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully fetched product", product))
}
