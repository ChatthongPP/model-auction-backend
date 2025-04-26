package controllers

import (
	"net/http"
	"strconv"

	"backend-service/internal/domain"
	"backend-service/pkg/utilities/responses"
	"backend-service/pkg/utilities/validator"
	"github.com/labstack/echo/v4"
)

func (h *Controller) CreateAdmin(c echo.Context) error {
	var adminDTO AdminManagementDTO
	if err := c.Bind(&adminDTO); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	// Validate adminDTO
	if err := validator.Validate(adminDTO); err != nil {
		validationErrors := validator.FormatValidationErrors(err)
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	// Map DTO to domain model
	admin := &domain.AdminManagement{
		FirstName: adminDTO.FirstName,
		LastName:  adminDTO.LastName,
		Email:     adminDTO.Email,
		Tel:       adminDTO.Tel,
		Password:  adminDTO.Password,
		CompanyID: adminDTO.CompanyId,
		CreatedAt: currentTime(),
	}

	// Call the usecase to add the admin
	id, err := h.uc.CreateAdmin(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	adminId := map[string]interface{}{
		"Id": id,
	}

	// Return the response with the created admin ID
	return c.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Admin added successfully", adminId))
}

func (uc *Controller) GetAdmins(c echo.Context) error {
	// Extract limit and page from query parameters
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10 // Set a default limit if not provided or invalid
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1 // Set a default page if not provided or invalid
	}
	search := c.QueryParam("search")
	beginDate := c.QueryParam("startTime")
	endDate := c.QueryParam("endTime")
	// Call the usecase to get admins
	admins, totalCount, totalPages, err := uc.uc.GetAdminAll(limit, page, search, beginDate, endDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	// Convert admins to AdminManagementDTO
	var adminDTOs []*AdminManagementDTO
	for _, admin := range admins {
		var createAtTimeStr string
		if admin.CreatedAt != nil {
			createAtTimeStr = admin.CreatedAt.Format("2006-01-02 15:04:05")
		}
		adminDTO := &AdminManagementDTO{
			Id:        admin.Id,
			FirstName: admin.FirstName,
			LastName:  admin.LastName,
			Email:     admin.Email,
			Tel:       admin.Tel,
			CompanyId: admin.CompanyID,
			CompanyName: Company{
				Id:          admin.CompanyName.Id,
				CompanyName: admin.CompanyName.CompanyName,
			},
			CreatedAt: createAtTimeStr,
		}
		adminDTOs = append(adminDTOs, adminDTO)
	}
	response := map[string]interface{}{
		"currentPage": page,
		"totalCount":  totalCount,
		"totalPages":  totalPages,
		"records":     adminDTOs,
	}

	// Send the response back to the client
	return c.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully retrieved admins", response))
}

func (h *Controller) UpdateAdmin(c echo.Context) error {
	idStr := c.Param("admins-id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}
	var adminDTO AdminManagementDTO
	if err := c.Bind(&adminDTO); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	admins := &domain.AdminManagement{
		Id:        id,
		FirstName: adminDTO.FirstName,
		LastName:  adminDTO.LastName,
		Email:     adminDTO.Email,
		Tel:       adminDTO.Tel,
		CompanyID: adminDTO.CompanyId,
	}
	if err := h.uc.UpdateAdmin(admins); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully updated admin", nil))
}

func (h *Controller) DeleteAdmin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("admins-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	admin := domain.AdminManagement{
		Id: id,
	}

	if err := h.uc.DeleteAdmin(&admin); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Admin deleted successfully"})
}

func (h *Controller) GetAdminById(c echo.Context) error {
	// Extract admin ID from path parameter
	id, err := strconv.Atoi(c.Param("admin-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	// Call the usecase to get admin by ID
	admin, err := h.uc.GetAdminById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	if admin == nil {
		return c.JSON(http.StatusNotFound, responses.Error(http.StatusNotFound, "Admin not found"))
	}

	// Convert admin to DTO
	adminDTO := &AdminManagementDTO{
		Id:        admin.Id,
		FirstName: admin.FirstName,
		LastName:  admin.LastName,
		Email:     admin.Email,
		Tel:       admin.Tel,
	}

	return c.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully retrieved admin", adminDTO))
}

func (h *Controller) SuperAdminUpdatePassword(c echo.Context) error {
	// ดึงค่า ID จากพารามิเตอร์
	idStr := c.Param("admins-id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Invalid admin ID"))
	}

	// ผูกข้อมูลจาก request body ไปยัง adminDTO
	var adminDTO AdminManagementDTO
	if err := c.Bind(&adminDTO); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	// ตรวจสอบว่ามีการใส่รหัสผ่านหรือไม่
	if adminDTO.Password == "" {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Missing required fields: password"))
	}

	// สร้างโครงสร้าง admin จากข้อมูลที่ได้รับ
	admins := &domain.AdminManagement{
		Id:       id,
		Password: adminDTO.Password,
	}

	// เรียกใช้ฟังก์ชันเพื่อเปลี่ยนรหัสผ่าน
	if err := h.uc.SuperAdminChangePassword(admins); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	}

	// ส่งคำตอบสำเร็จ
	return c.JSON(http.StatusOK, responses.Ok(http.StatusOK, "Successfully changed password", nil))
}
