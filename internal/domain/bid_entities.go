package domain

import "time"

type Bid struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	UserID    int       `json:"user_id"`
	BidAmount float64   `json:"bid_amount"`
	BidTime   time.Time `json:"bid_time"`
}
