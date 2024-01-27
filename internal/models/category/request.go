package category

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name"`
}

type GetAllCategoryRequest struct {
	CategoryName *string `query:"category_name"`
	Search       *string `query:"search"`
	Page         int     `query:"page"`
	PageSize     int     `query:"page_size"`
	OrderBy      string  `query:"order_by"`
	OrderType    string  `query:"order_type"`
}

type UpdateCategoryRequest struct {
	Id           int64  `json:"id"`
	CategoryName string `json:"category_name"`
}
