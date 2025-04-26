package domain

import "time"

type AdminManagement struct {
	Id          int        `json:"id"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Email        string     `json:"email"`
	Password    string     `json:"password"`
	Tel         string     `json:"tel"`
	CompanyID   int        `json:"companyId"`
	CompanyName Company    `json:"companyName"`
	DeletedAt   *time.Time `json:"deleteAt"`
	CreatedAt   *time.Time `json:"createAt"`
}
