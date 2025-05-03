package controllers

import "time"

type BidRequest struct {
	ID        int        `json:"id" validate:"omitempty"`
	ProductID int        `json:"product_id" validate:"omitempty"`
	UserID    int        `json:"user_id" validate:"required"`
	BidAmount float64    `json:"bid_amount" validate:"required,gt=0"`
	BidTime   time.Time  `json:"bid_time" validate:"omitempty"`
	CreatedAt time.Time  `json:"created_at" validate:"omitempty"`
	UpdatedAt time.Time  `json:"updated_at" validate:"omitempty"`
	DeletedAt *time.Time `json:"deleted_at" validate:"omitempty"`
}

type BidResponse struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	UserID    int       `json:"user_id"`
	BidAmount float64   `json:"bid_amount"`
	BidTime   time.Time `json:"bid_time"`
}
