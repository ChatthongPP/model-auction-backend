package domain

type Line struct {
	Id          int        `json:"id"`
	UID         string     `json:"uid"`
	DisplayName string     `json:"displayName"`
	PictureURL  string     `json:"pictureUrl"`
}
