package repositories

import (

	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
func (r *Repo) GetAdminByEmailAndPassword(admin *domain.Login) error {
	var admins models.LoginModel
	// ค้นหาผู้ใช้ด้วยอีเมล
	if err := r.db.Where("email = ? AND is_delete is NULL", admin.Email).First(&admins).Error; err != nil {
		return err // หากไม่พบผู้ใช้ ส่ง error กลับ
	}

	// เปรียบเทียบรหัสผ่าน
	if err := bcrypt.CompareHashAndPassword([]byte(admins.Password), []byte(admin.Password)); err != nil {
		return err // หากรหัสผ่านไม่ตรงกัน ส่ง error กลับ
	}

	// สร้าง struct ผู้ใช้
	admin.Id = admins.Id
	admin.Email = admins.Email

	return nil
}
func (r *Repo) GetById(adminId int) (*domain.Login, error) {
	var dbAdmin models.LoginModel

	if err := r.db.Where("id = ? AND is_delete is NULL", adminId).First(&dbAdmin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	admin := &domain.Login{
		Id:          dbAdmin.Id,
		Email:        dbAdmin.Email,
		Password:    dbAdmin.Password,
	}

	return admin, nil
}
