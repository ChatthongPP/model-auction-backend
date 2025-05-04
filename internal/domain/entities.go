package domain

type FilterRequest struct {
	CurrrentPage int    `json:"current_page"`
	Limit        int    `json:"limit"`
	OrderBy      string `json:"order_by"`
	Order        string `json:"order"`
	Search       string `json:"search"`
	CategoryID   int    `json:"category_id"`
	Status       string `json:"status"`
	SellerID     int    `json:"seller_id"`
	UserID       int    `json:"user_id"`
	ProductID    int    `json:"product_id"`
}

type PaginationResponse struct {
	CurrrentPage int         `json:"currrent_page"`
	Limit        int         `json:"limit"`
	TotalCount   int         `json:"total_count"`
	TotalPages   int         `json:"total_pages"`
	Data         interface{} `json:"data"`
}
