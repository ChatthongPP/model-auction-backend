package controllers


type AdminManagementDTO struct {
	Id        int        `json:"id"`
	FirstName string     `json:"firstName" validate:"required,max=32,isThaiOrEnglish"`
	LastName  string     `json:"lastName" validate:"required,max=32,isThaiOrEnglish"`
	Email      string     `json:"email" validate:"required,email,isASCII"`
	Tel       string     `json:"tel" validate:"required,len=10,numeric"`
	Password  string     `json:"password" validate:"required,min=8,max=32,isComplexPassword"`
	CompanyId int        `json:"companyId" validate:"required"`
	CompanyName Company    `json:"companyName"`
	CreatedAt  string `json:"createdAt"`
}
