package domain

type BidInterface interface {
	CreateBid(bid *Bid) error
	GetBids(filter *FilterRequest, offset int) ([]*Bid, int, error)
}
