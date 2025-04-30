package domain

import "time"

type Product struct {
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	CategoryID          int       `json:"category_id"`
	SellerID            int       `json:"seller_id"`
	ActualPrice         float64   `json:"actual_price"`
	StartingBidPrice    float64   `json:"starting_bid_price"`
	CurrentBidPrice     float64   `json:"current_bid_price"`
	MinimumBidIncrement float64   `json:"minimum_bid_increment"`
	ShippingPrice       float64   `json:"shipping_price"`
	ServiceFee          float64   `json:"service_fee"`
	AuctionStartTime    time.Time `json:"auction_start_time"`
	AuctionEndTime      time.Time `json:"auction_end_time"`
	Status              string    `json:"status"`
	// Image            string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
