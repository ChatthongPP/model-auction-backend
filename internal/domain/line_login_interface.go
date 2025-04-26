package domain

type LineInterface interface {
	CreateLineUser(lineUser *Line) error
	FindLineUserByUID(lineUID string) (*Line, error)
	VerifyIDToken(idToken string) (*Line, error)
	UpdateLineUser(lineUser Line) error
}