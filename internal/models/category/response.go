package category

import (
	"warabiz/api/pkg/infra/db"
)

type CreateCategoryResponse struct {
	Id int64 `json:"id"`
}

type GetAllCategoryResponse struct {
	Category   []CategoryList         `json:"warabisnis"`
	Pagination *db.PaginationResponse `json:"pagination"`
}

type CategoryDetailResponse struct {
	Category []Category `json:"warabisinis"`
}
