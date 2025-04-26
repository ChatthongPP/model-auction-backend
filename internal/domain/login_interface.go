package domain

type LoginInterface interface {
	GetById(adminId int) (*Login, error) 
	GetAdminByEmailAndPassword(admin *Login) error
}
