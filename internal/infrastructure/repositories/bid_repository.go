package repositories

import (
	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"

	lop "github.com/samber/lo/parallel"
)

func (r *Repo) CreateBid(bid *domain.Bid) error {
	dbBid := &models.BidModel{
		ProductID: bid.ProductID,
		UserID:    bid.UserID,
		BidAmount: bid.BidAmount,
		BidTime:   bid.BidTime,
		CreatedAt: bid.CreatedAt,
		UpdatedAt: bid.UpdatedAt,
	}

	err := r.db.Create(dbBid).Error
	if err != nil {
		return err
	}

	bid.ID = dbBid.ID

	return nil
}

func (r *Repo) GetBids(filter *domain.FilterRequest, offset int) ([]*domain.Bid, int, error) {
	var bids []*models.BidModel
	var totalCount int64

	query := r.db.Model(&models.BidModel{})

	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.ProductID != 0 {
		query = query.Where("product_id = ?", filter.ProductID)
	}
	if filter.OrderBy != "" {
		query = query.Order(filter.OrderBy + " " + filter.Order)
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(filter.Limit).Find(&bids).Error; err != nil {
		return nil, 0, err
	}

	domainBids := lop.Map(bids, func(bid *models.BidModel, i int) *domain.Bid {
		return &domain.Bid{
			ID:        bid.ID,
			ProductID: bid.ProductID,
			UserID:    bid.UserID,
			BidAmount: bid.BidAmount,
			BidTime:   bid.BidTime,
		}
	})

	return domainBids, int(totalCount), nil
}
