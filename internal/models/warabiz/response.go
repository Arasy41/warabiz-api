package warabiz

import (
	"warabiz/api/pkg/infra/db"
)

type CreateWarabizResponse struct {
	Id int64 `json:"id"`
}

type GetAllWarabizResponse struct {
	Warabiz    []WarabizList          `json:"warabisnis"`
	Pagination *db.PaginationResponse `json:"pagination"`
}

type WarabizDetailResponse struct {
	Warabiz      []Warabiz `json:"warabisnis"`
	CategoryName string    `json:"category_name"`
}
