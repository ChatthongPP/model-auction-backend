package controllers

type LoginLineRequest struct {
	IDToken   string     `json:"idToken"`
	LineUID   string     `json:"lineUid"`
}