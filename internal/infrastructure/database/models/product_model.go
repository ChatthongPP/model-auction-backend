package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductModel struct {
	ID                  int       `gorm:"column:id;primary_key" json:"id"`
	Name                string    `gorm:"column:name" json:"name"`
	Description         string    `gorm:"column:description" json:"description"`
	CategoryID          int       `gorm:"column:category_id" json:"category_id"`
	SellerID            int       `gorm:"column:seller_id" json:"seller_id"`
	ActualPrice         float64   `gorm:"column:actual_price" json:"actual_price"`
	StartingBidPrice    float64   `gorm:"column:starting_bid_price" json:"starting_bid_price"`
	CurrentBidPrice     float64   `gorm:"column:current_bid_price" json:"current_bid_price"`
	MinimumBidIncrement float64   `gorm:"column:minimum_bid_increment" json:"minimum_bid_increment"`
	ShippingPrice       float64   `gorm:"column:shipping_price" json:"shipping_price"`
	ServiceFee          float64   `gorm:"column:service_fee" json:"service_fee"`
	AuctionStartTime    time.Time `gorm:"column:auction_start_time" json:"auction_start_time"`
	AuctionEndTime      time.Time `gorm:"column:auction_end_time" json:"auction_end_time"`
	Status              string    `gorm:"column:status" json:"status"`
	// Image               image.Image    `gorm:"column:image" json:"image"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`

	Category CategoryModel `gorm:"foreignKey:CategoryID"`
	User     User          `gorm:"foreignKey:SellerID"`
}

func (*ProductModel) TableName() string {
	return "products"
}
