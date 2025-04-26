package controllers

import "time"

type StudentDTO struct {
	CreatedAt *time.Time `json:"createdAt" example:"2023-05-17 23:50:50"`
	UpdatedAt *time.Time `json:"updatedAt" example:"2023-05-17 23:50:50"`
	FirstName string     `json:"firstName" example:"Jimmy"`
	LastName  string     `json:"lastName" example:"Karuture"`
	Email     string     `json:"email" example:"jimmy@hiso.com"`
}
