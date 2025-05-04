package controllers

import (
	"net/http"
	"strconv"
	"time"

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
		ServiceFee:          productReq.ServiceFee,
		Image:               productReq.Image,
		AuctionStartTime:    productReq.AuctionStartTime,
		AuctionEndTime:      productReq.AuctionEndTime,
		Status:              productReq.Status,
		CreatedAt:           time.Now().UTC(),
		UpdatedAt:           time.Now().UTC(),
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

func (h *Controller) GetProducts(ctx echo.Context) error {
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
		CategoryID:   filterReq.CategoryID,
		Status:       filterReq.Status,
	}

	products, totalCount, totalPages, err := h.uc.GetProducts(filter)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	pagination := PaginationResponse{
		CurrrentPage: filter.CurrrentPage,
		Limit:        filter.Limit,
		TotalCount:   totalCount,
		TotalPages:   totalPages,
		Data:         products,
	}

	return ctx.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully fetched products", pagination))
}

func (h *Controller) UpdateProduct(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	var productReq ProductRequest
	if err := c.Bind(&productReq); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	if err := validator.Validate(&productReq); err != nil {
		validationErrors := validator.FormatValidationErrors(err)
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	product := &domain.Product{
		ID:                  id,
		Name:                productReq.Name,
		Description:         productReq.Description,
		CategoryID:          productReq.CategoryID,
		SellerID:            productReq.SellerID,
		ActualPrice:         productReq.ActualPrice,
		StartingBidPrice:    productReq.StartingBidPrice,
		CurrentBidPrice:     productReq.StartingBidPrice,
		MinimumBidIncrement: productReq.MinimumBidIncrement,
		ShippingPrice:       productReq.ShippingPrice,
		ServiceFee:          productReq.ServiceFee,
		Image:               productReq.Image,
		AuctionStartTime:    productReq.AuctionStartTime,
		AuctionEndTime:      productReq.AuctionEndTime,
		Status:              productReq.Status,
		UpdatedAt:           time.Now().UTC(),
	}

	if err := h.uc.UpdateProduct(product); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully updated product", product))
}
