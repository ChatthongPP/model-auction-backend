package usecase

import (
	"errors"

	"backend-service/internal/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (uc *Usecase) GetAdminAll(limit, page int, search, beginDate, endDate string) ([]*domain.AdminManagement, int,int, error) {
	offset := (page - 1) * limit
	admins, totalCount, err := uc.manageAdmin.GetAdminAll(limit, offset, search,beginDate,endDate)
	if err != nil {
		return nil, 0,0, err
	}
	totalPages := (totalCount + limit - 1) / limit

	return admins, totalCount, totalPages,nil
}

func (uc *Usecase) UpdateAdmin(admin *domain.AdminManagement) error {
	// Add input validation for admin object if necessary

	// Call the repository to update the admin record
	err := uc.manageAdmin.UpdateAdmin(admin)
	if err != nil {
		return err
	}

	return nil
}

func (uc *Usecase) DeleteAdmin(admin *domain.AdminManagement) error {
	err := uc.manageAdmin.DeleteAdmin(admin)
	if err != nil {
		return err
	}

	return nil
}

func (uc *Usecase) CreateAdmin(admin *domain.AdminManagement) (int, error) {
	existingMail, err := uc.manageAdmin.GetEmail(admin.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	if err == nil && existingMail.Id != 0 {
		return 0, errors.New("บัญชี "+ admin.Email +" มีอยู่แล้ว"  )
	}

	// ทำการเข้ารหัสรหัสผ่านก่อนที่จะเพิ่ม Admin
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, errors.New("error hashing password: " + err.Error())
	}

	// แทนที่รหัสผ่านเดิมด้วยรหัสผ่านที่เข้ารหัสแล้ว
	admin.Password = string(hashedPassword)

	// เรียกฟังก์ชัน AddAdmin ของ repository และรับค่า id ที่คืนมา
	id, err := uc.manageAdmin.CreateAdmin(admin)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *Usecase) GetAdminById(adminId int) (*domain.AdminManagement, error) {
	admin, err := uc.manageAdmin.GetAdminById(adminId)
	if err != nil {
		return nil, err // คืน error ที่ได้รับจาก repository layer ไปที่ caller
	}

	return admin, nil // คืนข้อมูล admin ที่ได้รับจาก repository layer
}

func (uc *Usecase) SuperAdminChangePassword(admin *domain.AdminManagement) error {
	// แฮชรหัสผ่าน
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password: " + err.Error())
	}

	// แทนที่รหัสผ่านเดิมด้วยรหัสผ่านที่เข้ารหัสแล้ว
	admin.Password = string(hashedPassword)

	// เรียกใช้ repository เพื่ออัปเดตข้อมูลของ admin
	if err := uc.manageAdmin.SuperAdminChangePassword(admin); err != nil {
		return err
	}

	return nil
}