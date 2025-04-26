package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"backend-service/internal/domain"
	"backend-service/pkg/utilities/responses"

	"github.com/labstack/echo/v4"
)

func (h *Controller) GetCompany(c echo.Context) error {
	companyList, err := h.uc.GetCompany()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// สร้าง slice เพื่อเก็บข้อมูลที่แปลงแล้ว
	var companyResponses []Company

	// วนลูปเพื่อแปลงข้อมูลแต่ละรายการใน companyList
	for _, company := range companyList {

		companyResponse := Company{
			Id:          company.Id,
			CompanyName: company.CompanyName,
		}
		// เพิ่มข้อมูลลงใน slice companyResponses
		companyResponses = append(companyResponses, companyResponse)
	}

	return c.JSON(http.StatusOK, companyResponses)
}

func (h *Controller) AddCompany(c echo.Context) error {
	var addCompany Company
	if err := c.Bind(&addCompany); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}
	if addCompany.CompanyName == "" {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Missing required fields"))
	}
	company := domain.Company{
		CompanyName: addCompany.CompanyName,
	}

	err := h.uc.AddCompany(company)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	} else {
		return c.JSON(http.StatusCreated, responses.Ok(http.StatusCreated, "Successfully save", nil))
	}
}

func (h *Controller) DeleteCompanyById(c echo.Context) error {
	// Extract company ID from path parameter
	id, err := strconv.Atoi(c.Param("company-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid company ID"})
	}

	// Create a Company struct with ID for repository function
	company := domain.Company{
		Id: id,
	}

	// Call the usecase to delete the company by ID (soft delete)
	if err := h.uc.DeleteCompanyById(company); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Failed to delete company: %s", err.Error())})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Company deleted successfully"})
}

func (h *Controller) UpdateCompany(c echo.Context) error {
	idStr := c.Param("company-id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest,err.Error()))
	}

	var updateCompany domain.Company
	if err := c.Bind(&updateCompany); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest,err.Error()))
	}

	if updateCompany.CompanyName == "" {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Company name cannot be empty"))
	}

	company := domain.Company{
		Id:          id,
		CompanyName: updateCompany.CompanyName,
	}
	if err := h.uc.UpdateCompany(company); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully updated company", nil))
}
