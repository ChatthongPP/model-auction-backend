package controllers

type Login struct {
	Id       int    `json:"id"`
	Email    string `json:"email" validate:"required,email,isASCII"`
	Password string `json:"password" validate:"required"`
}
