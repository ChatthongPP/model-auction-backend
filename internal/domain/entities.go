package domain

type FilterRequest struct {
	CurrrentPage int    `json:"current_page"` // หน้าปัจจุบัน
	Limit        int    `json:"limit"`        // จำนวนรายการต่อหน้า
	OrderBy      string `json:"order_by"`     // ชื่อคอลัมน์สำหรับจัดเรียง
	Order        string `json:"order"`        // ทิศทางการจัดเรียง (asc/desc)
	Search       string `json:"search"`       // คำค้นหา
	CategoryID   int    `json:"category_id"`  // กรองตาม Category ID
	Status       string `json:"status"`       // กรองตามสถานะ
}

type PaginationResponse struct {
	CurrrentPage int         `json:"currrent_page"` // หน้าปัจจุบัน
	Limit        int         `json:"limit"`         // จำนวนรายการต่อหน้า
	TotalCount   int         `json:"total_count"`   // จำนวนรายการทั้งหมด
	TotalPages   int         `json:"total_pages"`   // จำนวนหน้าทั้งหมด
	Data         interface{} `json:"data"`          // ข้อมูลรายการ
}
