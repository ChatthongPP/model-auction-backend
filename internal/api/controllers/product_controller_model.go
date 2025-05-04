package controllers

import "time"

type ProductRequest struct {
	Name                string    `json:"name" validate:"required,min=1,max=100"`
	Description         string    `json:"description" validate:"omitempty,min=1"`
	CategoryID          int       `json:"category_id" validate:"required"`
	SellerID            int       `json:"seller_id" validate:"required"`
	ActualPrice         float64   `json:"actual_price" validate:"omitempty,gt=0"`
	StartingBidPrice    float64   `json:"starting_bid_price" validate:"omitempty,gt=0"`
	MinimumBidIncrement float64   `json:"minimum_bid_increment" validate:"omitempty,gt=0"`
	ShippingPrice       float64   `json:"shipping_price" validate:"gte=0"`
	ServiceFee          float64   `json:"service_fee" validate:"gte=0"`
	AuctionStartTime    time.Time `json:"auction_start_time" validate:"omitempty"`
	AuctionEndTime      time.Time `json:"auction_end_time" validate:"omitempty,gtfield=AuctionStartTime"`
	Status              string    `json:"status" validate:"omitempty,oneof=active inactive"`
	Image               []string  `json:"image" validate:"omitempty"`
}
