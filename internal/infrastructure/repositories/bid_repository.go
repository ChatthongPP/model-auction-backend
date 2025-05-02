package repositories

import (
	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"
)

func (r *Repo) CreateBid(bid *domain.Bid) error {
	dbBid := &models.BidModel{
		ProductID: bid.ProductID,
		UserID:    bid.UserID,
		BidAmount: bid.BidAmount,
		BidTime:   bid.BidTime,
	}

	err := r.db.Create(dbBid).Error
	if err != nil {
		return err
	}

	bid.ID = dbBid.ID

	return nil
}
