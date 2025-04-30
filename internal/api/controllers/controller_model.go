package controllers

type FilterRequest struct {
	CurrrentPage int    `json:"current_page" query:"current_page" validate:"required,min=1"` // หน้าปัจจุบัน
	Limit        int    `json:"limit" query:"limit" validate:"required,min=1"`               // จำนวนรายการต่อหน้า
	OrderBy      string `json:"order_by" query:"order_by" validate:"omitempty"`              // ชื่อคอลัมน์สำหรับจัดเรียง
	Order        string `json:"order" query:"order" validate:"required,oneof=asc desc"`      // ทิศทางการจัดเรียง (asc/desc)
	Search       string `json:"search" query:"search" validate:"omitempty"`                  // คำค้นหา
	CategoryID   int    `json:"category_id" query:"category_id" validate:"omitempty"`        // กรองตาม Category ID
	Status       string `json:"status" query:"status" validate:"omitempty"`                  // กรองตามสถานะ
}

type PaginationResponse struct {
	CurrrentPage int         `json:"currrent_page"` // หน้าปัจจุบัน
	Limit        int         `json:"limit"`         // จำนวนรายการต่อหน้า
	TotalCount   int         `json:"total_count"`   // จำนวนรายการทั้งหมด
	TotalPages   int         `json:"total_pages"`   // จำนวนหน้าทั้งหมด
	Data         interface{} `json:"data"`          // ข้อมูลรายการ
}
