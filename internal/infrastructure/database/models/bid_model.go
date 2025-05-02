package models

import "time"

type BidModel struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	ProductID int       `gorm:"column:product_id" json:"product_id"`
	UserID    int       `gorm:"column:user_id" json:"user_id"`
	BidAmount float64   `gorm:"column:bid_amount" json:"bid_amount"`
	BidTime   time.Time `gorm:"column:bid_time" json:"bid_time"`
}

func (*BidModel) TableName() string {
	return "bids"
}
