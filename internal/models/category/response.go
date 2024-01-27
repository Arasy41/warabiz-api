package category

import (
	"warabiz/api/internal/models/warabiz"
	"warabiz/api/pkg/infra/db"
)

type CreateCategoryResponse struct {
	Id int64 `json:"id"`
}

type GetAllCategoryResponse struct {
	Category   []CategoryList         `json:"category"`
	Pagination *db.PaginationResponse `json:"pagination"`
}

type CategoryDetailResponse struct {
	Category []Category `json:"warabisinis"`
}

type GetCategoryDetail struct {
	Id           int64             `json:"id"`
	WarabizId    int64             `json:"warabiz_id"`
	Warabiz      []warabiz.Warabiz `json:"warabiz"`
	CategoriesId int64             `json:"categories_id"`
	Category     []Category        `json:"categories"`
}
