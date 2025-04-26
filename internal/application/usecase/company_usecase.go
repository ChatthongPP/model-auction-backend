package usecase

import (
	"errors"

	"backend-service/internal/domain"

	"gorm.io/gorm"
)

func (uc *Usecase) GetCompany() ([]domain.Company, error) {
	companies, err := uc.companyRepo.GetAllCompanies()
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (uc *Usecase) AddCompany(company domain.Company) error {
	// ตรวจสอบว่ามีบริษัทที่มีชื่อเดียวกันอยู่แล้วหรือไม่
	existingCompany, err := uc.companyRepo.GetCompanyByName(company.CompanyName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err == nil && existingCompany.Id != 0 {
		return errors.New("บริษัทที่มีชื่อ %s มีอยู่แล้ว" + company.CompanyName)
	}

	// เพิ่มบริษัทใหม่ลงในฐานข้อมูล
	return uc.companyRepo.AddCompany(company)
}

func (uc *Usecase) DeleteCompanyById(company domain.Company) error {
	err := uc.companyRepo.DeleteCompanyById(company)
	if err != nil {
		return err
	}

	return nil
}

func (uc *Usecase) UpdateCompany(company domain.Company) error {
	// ตรวจสอบว่ามีบริษัทที่มีชื่อเดียวกันอยู่แล้วหรือไม่
	existingCompany, err := uc.companyRepo.GetCompanyByName(company.CompanyName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err == nil {
		if existingCompany.Id != company.Id {
			return errors.New("บริษัทที่มีชื่อ " + company.CompanyName + " มีอยู่แล้ว")
		}
	}
	uc.companyRepo.UpdateCompany(company)
	return nil
}
