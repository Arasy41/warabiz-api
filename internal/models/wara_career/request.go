package wara_career

import "mime/multipart"

type CreateWaraCareerRequest struct {
	Id          int64                 `json:"id" validate:"required"`
	CareerTitle string                `json:"career_title" validate:"required"`
	Description string                `json:"description" validate:"required"`
	Address     string                `json:"address" validate:"required"`
	ImageUrl    *multipart.FileHeader `json:"image_url"`
}

type GetAllWaraCareerRequest struct {
	Search    *string `query:"search"`
	Page      int     `query:"page"`
	PageSize  int     `query:"page_size"`
	OrderBy   string  `query:"order_by"`
	OrderType string  `query:"order_type"`
}

type UpdateWaraCareerRequest struct {
	Id          int64                 `json:"id"`
	CareerTitle string                `json:"career_title"`
	Description string                `json:"description"`
	Address     string                `json:"address"`
	ImageUrl    *multipart.FileHeader `json:"image_url"`
}
