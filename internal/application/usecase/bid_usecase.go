package usecase

import (
	"time"

	"backend-service/internal/domain"
)

func (uc *Usecase) CreateBid(bidReq *domain.Bid) error {
	bid := &domain.Bid{
		ProductID: bidReq.ProductID,
		UserID:    bidReq.UserID,
		BidAmount: bidReq.BidAmount,
		BidTime:   bidReq.BidTime,
	}

	product, err := uc.productRepo.GetProductByID(bidReq.ProductID)
	if err != nil {
		return err
	}

	currentTime := time.Now()
	timeRemaining := product.AuctionEndTime.Sub(currentTime)
	if timeRemaining <= time.Minute {
		newEndTime := currentTime.Add(time.Minute)
		product.AuctionEndTime = newEndTime

		err = uc.productRepo.UpdateProduct(product)
		if err != nil {
			return err
		}
	}

	if err := uc.bidRepo.CreateBid(bid); err != nil {
		return err
	}

	return nil
}
