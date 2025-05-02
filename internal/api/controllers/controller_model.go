package controllers

type FilterRequest struct {
	CurrrentPage int    `json:"current_page" query:"current_page" validate:"required,min=1"`
	Limit        int    `json:"limit" query:"limit" validate:"required,min=1"`
	OrderBy      string `json:"order_by" query:"order_by" validate:"omitempty"`
	Order        string `json:"order" query:"order" validate:"required,oneof=asc desc"`
	Search       string `json:"search" query:"search" validate:"omitempty"`
	CategoryID   int    `json:"category_id" query:"category_id" validate:"omitempty"`
	Status       string `json:"status" query:"status" validate:"omitempty"`
}

type PaginationResponse struct {
	CurrrentPage int         `json:"currrent_page"`
	Limit        int         `json:"limit"`
	TotalCount   int         `json:"total_count"`
	TotalPages   int         `json:"total_pages"`
	Data         interface{} `json:"data"`
}
