package domain

import "time"

type Company struct {
	Id          int    `json:"id"`
	CompanyName string `json:"companyName"`
	DeletedAt   *time.Time `json:"deleteAt"`
}