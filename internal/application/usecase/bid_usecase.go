package usecase

import (
	"time"

	"backend-service/internal/domain"
)

func (uc *Usecase) CreateBid(bid *domain.Bid) error {
	product, err := uc.productRepo.GetProductByID(bid.ProductID)
	if err != nil {
		return err
	}

	product.CurrentBidPrice = bid.BidAmount
	currentTime := bid.BidTime
	timeRemaining := product.AuctionEndTime.Sub(currentTime)
	if timeRemaining <= time.Minute {
		newEndTime := product.AuctionEndTime.Add(time.Minute)
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

func (uc *Usecase) GetBids(filter *domain.FilterRequest) ([]*domain.Bid, int, int, error) {
	offset := (filter.CurrrentPage - 1) * filter.Limit

	bids, totalCount, err := uc.bidRepo.GetBids(filter, offset)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := (totalCount + filter.Limit - 1) / filter.Limit

	return bids, totalCount, totalPages, nil
}
