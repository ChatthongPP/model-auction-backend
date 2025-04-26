package domain

type AdminManagementInterface interface {
	CreateAdmin(admin *AdminManagement) (int, error) 
	GetAdminById(adminId int) (*AdminManagement, error) 
	GetAdminAll(int, int, string, string, string) ([]*AdminManagement, int, error)
	DeleteAdmin(admin *AdminManagement) error
	UpdateAdmin(admin *AdminManagement) error
	SuperAdminChangePassword(admin *AdminManagement) error
	GetEmail(mail string) (*AdminManagement, error)
}
