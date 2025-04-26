package domain

import "time"

type Login struct {
	Id          int        `json:"id"`
	Email        string     `json:"email"`
	Password    string     `json:"password"`
	DeletedAt   *time.Time `json:"deleteAt"`
	CreateAt    *time.Time `json:"createAt"`
}
