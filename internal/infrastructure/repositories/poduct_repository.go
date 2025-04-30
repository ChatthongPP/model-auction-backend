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

func (r *Repo) GetProductByID(id int) (*domain.Product, error) {
	dbProduct := &models.ProductModel{}
	err := r.db.First(dbProduct, id).Error
	if err != nil {
		return nil, err
	}

	product := &domain.Product{
		ID:                  dbProduct.ID,
		Name:                dbProduct.Name,
		Description:         dbProduct.Description,
		CategoryID:          dbProduct.CategoryID,
		SellerID:            dbProduct.SellerID,
		ActualPrice:         dbProduct.ActualPrice,
		StartingBidPrice:    dbProduct.StartingBidPrice,
		CurrentBidPrice:     dbProduct.CurrentBidPrice,
		MinimumBidIncrement: dbProduct.MinimumBidIncrement,
		ShippingPrice:       dbProduct.ShippingPrice,
		ServiceFee:          dbProduct.ServiceFee,
		AuctionStartTime:    dbProduct.AuctionStartTime,
		AuctionEndTime:      dbProduct.AuctionEndTime,
		Status:              dbProduct.Status,
		// Image:            dbProduct.Image,
		CreatedAt: dbProduct.CreatedAt,
		UpdatedAt: dbProduct.UpdatedAt,
	}

	return product, nil
}
