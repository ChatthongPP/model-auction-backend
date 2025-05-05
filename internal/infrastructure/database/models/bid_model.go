package models

import "time"

type BidModel struct {
	ID        int        `gorm:"column:id;primary_key" json:"id"`
	ProductID int        `gorm:"column:product_id" json:"product_id"`
	UserID    int        `gorm:"column:user_id" json:"user_id"`
	BidAmount float64    `gorm:"column:bid_amount" json:"bid_amount"`
	BidTime   time.Time  `gorm:"column:bid_time" json:"bid_time"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	User      User       `gorm:"foreignKey:UserID"`
	Product   Product    `gorm:"foreignKey:ProductID"`
}

func (*BidModel) TableName() string {
	return "bids"
}
