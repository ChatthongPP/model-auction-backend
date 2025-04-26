package repositories

import (
	"time"

	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

// GetAdminById retrieves admin information by ID.
func (r *Repo) GetAdminById(adminId int) (*domain.AdminManagement, error) {
	var dbAdmin models.AdminModel

	if err := r.db.Preload("email").Preload("password").Where("id = ?", adminId).First(&dbAdmin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	admin := &domain.AdminManagement{
		Id:          dbAdmin.Id,
		FirstName:   dbAdmin.FirstName,
		LastName:    dbAdmin.LastName,
		Email:       dbAdmin.Email.Email,
		Password:    dbAdmin.Password.Password,
		Tel:         dbAdmin.Tel,
		CompanyID:   dbAdmin.CompanyID,
		CompanyName: domain.Company{Id: dbAdmin.CompanyName.Id, CompanyName: dbAdmin.CompanyName.CompanyName},
	}

	return admin, nil
}

// GetAdminAll retrieves all admins with pagination and search.
func (r *Repo) GetAdminAll(limit, offset int, search, beginDate, endDate string) ([]*domain.AdminManagement, int, error) {
	var dbAdmins []models.AdminModel
	var totalCount int64

	// Build the filter query
	conditionStr, args, err := r.buildFilterQuery("admin_info", search, beginDate, endDate)
	if err != nil {
		return nil, 0, err
	}

	// Count the total number of records that match the filter
	err = r.db.Model(&models.AdminModel{}).
		Joins("LEFT JOIN company ON company.id = admin_info.company_id").
		Joins("LEFT JOIN login ON login.id  = admin_info.mail_id").
		Where(conditionStr, args...).
		Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	// Retrieve the records that match the filter
	err = r.db.Preload("Email").Preload("Password").Preload("CompanyName").
		Joins("LEFT JOIN company ON company.id = admin_info.company_id").
		Joins("LEFT JOIN login ON login.id  = admin_info.mail_id").
		Where(conditionStr, args...).
		Order("admin_info.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&dbAdmins).Error
	if err != nil {
		return nil, 0, err
	}

	// Convert results to domain.AdminManagement
	admins := make([]*domain.AdminManagement, len(dbAdmins))
	for i, dbAdmin := range dbAdmins {
		admins[i] = &domain.AdminManagement{
			Id:          dbAdmin.Id,
			FirstName:   dbAdmin.FirstName,
			LastName:    dbAdmin.LastName,
			Email:       dbAdmin.Email.Email,
			Password:    dbAdmin.Password.Password,
			Tel:         dbAdmin.Tel,
			CompanyID:   dbAdmin.CompanyID,
			CompanyName: domain.Company{Id: dbAdmin.CompanyName.Id, CompanyName: dbAdmin.CompanyName.CompanyName},
			CreatedAt:   dbAdmin.CreatedAt,
		}
	}

	return admins, int(totalCount), nil
}

// UpdateAdmin updates admin information.
func (r *Repo) UpdateAdmin(admin *domain.AdminManagement) error {
	dbAdmin := models.AdminModel{
		FirstName: admin.FirstName,
		LastName:  admin.LastName,
		Tel:       admin.Tel,
		CompanyID: admin.CompanyID,
	}

	if err := r.db.Model(&models.AdminModel{}).Where("id=?", admin.Id).Updates(dbAdmin).Error; err != nil {
		return err
	}

	// Update LoginModel fields if necessary
	if admin.Email != "" {
		if err := r.db.Model(&models.LoginModel{}).Where("id=?", dbAdmin.MailID).Update("email", admin.Email).Error; err != nil {
			return err
		}
	}
	if admin.Password != "" {
		if err := r.db.Model(&models.LoginModel{}).Where("id=?", dbAdmin.PasswordID).Update("password", admin.Password).Error; err != nil {
			return err
		}
	}

	return nil
}

// DeleteAdmin marks an admin as deleted.
func (r *Repo) DeleteAdmin(admin *domain.AdminManagement) error {
	now := time.Now()
	deleteAdmin := models.AdminModel{
		DeletedAt: &now,
	}
	where := models.AdminModel{
		Id: admin.Id,
	}
	if err := r.db.Model(&where).Updates(deleteAdmin).Error; err != nil {
		return err
	}

	return nil
}

// CreateAdmin creates a new admin.
func (r *Repo) CreateAdmin(admin *domain.AdminManagement) (int, error) {
	// Create a new LoginModel for mail and password together
	login := models.LoginModel{
		Email:    admin.Email,
		Password: admin.Password,
	}

	if err := r.db.Create(&login).Error; err != nil {
		return 0, err
	}

	dbAdmin := models.AdminModel{
		FirstName:  admin.FirstName,
		LastName:   admin.LastName,
		MailID:     login.Id, // Use MailID
		PasswordID: login.Id, // Use PasswordID
		Tel:        admin.Tel,
		CompanyID:  admin.CompanyID,
		CreatedAt:  admin.CreatedAt,
	}

	if err := r.db.Create(&dbAdmin).Error; err != nil {
		return 0, err
	}
	return dbAdmin.Id, nil
}

// SuperAdminChangePassword changes the password of a super admin.
func (r *Repo) SuperAdminChangePassword(admin *domain.AdminManagement) error {
	dbAdmin := models.AdminModel{}
	if err := r.db.Where("id=?", admin.Id).First(&dbAdmin).Error; err != nil {
		return err
	}

	if err := r.db.Model(&models.LoginModel{}).Where("id=?", dbAdmin.PasswordID).Update("password", admin.Password).Error; err != nil {
		return err
	}

	return nil
}

// GetEmail retrieves admin information by email.
func (r *Repo) GetEmail(mail string) (*domain.AdminManagement, error) {
	var login models.LoginModel
	if err := r.db.Where("email = ?", mail).First(&login).Error; err != nil {
		return nil, err
	}

	var admin models.AdminModel
	if err := r.db.Where("mail_id = ?", login.Id).First(&admin).Error; err != nil {
		return nil, err
	}

	return &domain.AdminManagement{
		Id:    admin.Id,
		Email: login.Email,
	}, nil
}
