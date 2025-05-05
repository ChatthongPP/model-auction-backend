package repositories

import (
	"encoding/json"
	"fmt"
	"time"

	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"

	"gorm.io/datatypes"

	lop "github.com/samber/lo/parallel"
)

func (r *Repo) CreateProduct(product *domain.Product) (err error) {
	// log value product
	fmt.Println("product", product)

	imagesJSON := []byte("[]")
	if product.Image != nil {
		imagesJSON, err = json.Marshal(product.Image)
		if err != nil {
			return err
		}

		fmt.Println("imagesJSON", imagesJSON)
	}

	dbProduct := &models.Product{
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
		Image:               imagesJSON,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	fmt.Printf("dbProduct: %+v\n", dbProduct)

	if err := r.db.Create(dbProduct).Error; err != nil {
		return err
	}

	product.ID = dbProduct.ID

	return nil
}

func (r *Repo) GetProductByID(id int) (*domain.Product, error) {
	dbProduct := &models.Product{}

	err := r.db.Preload("Category").Preload("User").First(dbProduct, id).Error
	if err != nil {
		return nil, err
	}

	var images []string
	err = json.Unmarshal([]byte(dbProduct.Image), &images)
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
		// CategoryName:        dbProduct.Category.Name,
		// SellerName:          dbProduct.User.FirstName + " " + dbProduct.User.LastName,
		Image: images,
	}

	return product, nil
}

func (r *Repo) GetProducts(filter *domain.FilterRequest, offset int) ([]*domain.Product, int, error) {
	var products []*models.Product
	var totalCount int64

	query := r.db.Model(&models.Product{})

	if filter.Search != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Search+"%")
	}
	if filter.CategoryID > 0 {
		query = query.Where("category_id = ?", filter.CategoryID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.SellerID != 0 {
		query = query.Where("seller_id = ?", filter.SellerID)
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

	domainProducts := lop.Map(products, func(product *models.Product, i int) *domain.Product {
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
	imagesJSON, err := json.Marshal(product.Image)
	if err != nil {
		return err
	}

	dbProduct := models.Product{
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
		Image:               datatypes.JSON(imagesJSON),
	}

	if err := r.db.Model(&dbProduct).Where("id = ?", product.ID).Updates(&dbProduct).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteProduct(id int) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		return err
	}

	return nil
}
