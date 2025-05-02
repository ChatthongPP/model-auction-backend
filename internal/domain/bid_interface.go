package domain

type BidInterface interface {
	CreateBid(bid *Bid) error
}
