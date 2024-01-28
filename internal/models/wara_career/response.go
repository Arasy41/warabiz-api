package wara_career

import "warabiz/api/pkg/infra/db"

type CreateWaraCareerResponse struct {
	Id int64 `json:"id"`
}

type GetAllWaraCareerResponse struct {
	WaraCareer []WaraCareerList       `json:"category"`
	Pagination *db.PaginationResponse `json:"pagination"`
}

type WaraCareerDetailResponse struct {
	WaraCareer []WaraCareer `json:"wara_career"`
}
