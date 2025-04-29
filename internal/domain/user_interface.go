package domain

type UserInterface interface {
	CreateUser(req User) (*User, error)
	FindByEmail(email string) (*User, error)
	GetUserByID(userID uint) (*User, error)
}
