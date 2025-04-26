package repositories

import (
	"errors"
	"time"

	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"
)

func (r *Repo) GetAllCompanies() ([]domain.Company, error) {
	var companies []models.CompanyModel

	if err := r.db.Where("is_delete IS NULL").Find(&companies).Error; err != nil {
		return nil, err
	}

	var result []domain.Company

	for _, company := range companies {
		result = append(result, domain.Company{
			Id:          company.Id,
			CompanyName: company.CompanyName,
		})
	}

	return result, nil
}

func (r *Repo) GetComPany(company domain.Company) error {
	var companyModel models.CompanyModel
	if err := r.db.Where("company_name = ?", company.CompanyName).First(&companyModel).Error; err != nil {
		return err
	}
	if companyModel.Id == 0 {
		return errors.New("company not found")
	}
	return nil
}

func (r *Repo) AddCompany(company domain.Company) error {
	dbRepo := models.CompanyModel{
		Id:          company.Id,
		CompanyName: company.CompanyName,
	}

	if err := r.db.Create(&dbRepo).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteCompanyById(company domain.Company) error {
	now := time.Now()
	DeleteCompany := models.CompanyModel{
		DeletedAt: &now,
	}
	WHERE := models.CompanyModel{
		Id: company.Id,
	}
	if err := r.db.Model(&WHERE).Updates(DeleteCompany).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) UpdateCompany(company domain.Company) error {
	var dbCompany models.CompanyModel
	if err := r.db.First(&dbCompany, company.Id).Error; err != nil {
		return err
	}

	dbCompany.CompanyName = company.CompanyName

	if err := r.db.Save(&dbCompany).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetCompanyId(companyID int) error {
	var companyModel models.CompanyModel
	if err := r.db.First(&companyModel, companyID).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetCompanyByName(companyName string) (*domain.Company, error) {
	var company models.CompanyModel
	if err := r.db.Where("company_name = ? AND is_delete is NULL", companyName).First(&company).Error; err != nil {
		return nil, err
	}

	return &domain.Company{
		Id:          company.Id,
		CompanyName: company.CompanyName,
	}, nil
}
