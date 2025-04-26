package domain

type CompanyListInterface interface {
	GetAllCompanies() ([]Company, error)
	GetComPany(company Company) error
	AddCompany(company Company) error
	DeleteCompanyById(company Company) error
	UpdateCompany(company Company) error
	GetCompanyByName(companyName string) (*Company, error)
}