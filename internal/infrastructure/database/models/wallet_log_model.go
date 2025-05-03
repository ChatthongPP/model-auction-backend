package models

import "time"

type WalletLogModel struct {
	ID          int        `gorm:"column:id;primary_key" json:"id"`
	Amount      float64    `gorm:"column:amount" json:"amount"`
	UserID      int        `gorm:"column:user_id" json:"user_id"`
	Status      string     `gorm:"column:status" json:"status"`
	UpdatedByID int        `gorm:"column:updated_by_id" json:"updated_by_id"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (*WalletLogModel) TableName() string {
	return "wallet_log"
}
