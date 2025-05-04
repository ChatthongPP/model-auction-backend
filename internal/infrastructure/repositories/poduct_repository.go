package repositories

import (
	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"

	lop "github.com/samber/lo/parallel"
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
		CreatedAt:           product.CreatedAt,
		UpdatedAt:           product.UpdatedAt,
		Image:               product.Image,
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

	err := r.db.Preload("Category").Preload("User").First(dbProduct, id).Error
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
		CreatedAt:           dbProduct.CreatedAt,
		UpdatedAt:           dbProduct.UpdatedAt,
		CategoryName:        dbProduct.Category.Name,
		SellerName:          dbProduct.User.FirstName + " " + dbProduct.User.LastName,
		Image:               dbProduct.Image,
	}

	return product, nil
}

func (r *Repo) GetProducts(filter *domain.FilterRequest, offset int) ([]*domain.Product, int, error) {
	var products []*models.ProductModel
	var totalCount int64

	query := r.db.Model(&models.ProductModel{})

	if filter.Search != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Search+"%")
	}
	if filter.CategoryID > 0 {
		query = query.Where("category_id = ?", filter.CategoryID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.OrderBy != "" {
		query = query.Order(filter.OrderBy + " " + filter.Order)
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	err := query.Limit(filter.Limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	domainProducts := lop.Map(products, func(product *models.ProductModel, i int) *domain.Product {
		return &domain.Product{
			ID:                  product.ID,
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
			CreatedAt:           product.CreatedAt,
			UpdatedAt:           product.UpdatedAt,
		}
	})

	return domainProducts, int(totalCount), nil
}

func (r *Repo) UpdateProduct(product *domain.Product) error {
	dbProduct := models.ProductModel{
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
		UpdatedAt:           product.UpdatedAt,
		Image:               product.Image,
	}

	if err := r.db.Model(&dbProduct).Where("id = ?", product.ID).Updates(&dbProduct).Error; err != nil {
		return err
	}

	return nil
}
