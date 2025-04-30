package repositories

import (
	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"
)

func (r *Repo) CreateProduct(product *domain.Product) error {
	dbProduct := &models.ProductModel{
		Name:                product.Name,
		Description:         product.Description,
		CategoryID:          product.CategoryID,
		SellerID:            product.SellerID,
		ActualPrice:         product.ActualPrice,
		StartingBidPrice:    product.StartingBidPrice,
		CurrentBidPrice:     product.CurrentBidPrice,
		MinimumBidIncrement: product.MinimumBidIncrement,
		ShippingPrice:       product.ShippingPrice,
		ServiceFee:          product.ServiceFee,
		AuctionStartTime:    product.AuctionStartTime,
		AuctionEndTime:      product.AuctionEndTime,
		Status:              product.Status,
		// Image:           product.Image,
	}

	err := r.db.Create(dbProduct).Error
	if err != nil {
		return err
	}

	product.ID = dbProduct.ID

	return nil
}
