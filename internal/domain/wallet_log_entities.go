package domain

import "time"

type WalletLog struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	UserID      int       `json:"user_id"`
	Status      string    `json:"status"`
	UpdatedByID int       `json:"updated_by_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
