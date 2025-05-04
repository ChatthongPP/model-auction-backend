package domain

import "time"

type Bid struct {
	ID          int       `json:"id"`
	ProductID   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	UserID      int       `json:"user_id"`
	BidderName  string    `json:"bidder_name"`
	BidAmount   float64   `json:"bid_amount"`
	BidTime     time.Time `json:"bid_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
