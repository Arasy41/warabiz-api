package category

import (
	"time"
	"warabiz/api/pkg/utils/sanitizer"
)

type CreateCategoryRequest struct {
	Id              int64  `json:"id" validate:"required"`
	CategoryId      int64  `json:"category_id" validate:"required"`
	WaralabaName    string `json:"waralaba_name" validate:"required"`
	Prize           string `json:"prize" validate:"required"`
	Contact         string `json:"contact" validate:"required"`
	BrochureLink    string `json:"brochure_link" validate:"required"`
	Since           string `json:"since" validate:"required"`
	OutletTotal     int64  `json:"outlet_total" validate:"required"`
	LicenseDuration string `json:"license_duration" validate:"required"`
}

func (m *CreateCategoryRequest) Sanitize() {
	m.WaralabaName = sanitizer.SanitizeHTML(m.WaralabaName)
	m.Prize = sanitizer.SanitizeHTML(m.Prize)
	m.Contact = sanitizer.SanitizeHTML(m.Contact)
	m.BrochureLink = sanitizer.SanitizeHTML(m.BrochureLink)
}

type GetAllCategoryRequest struct {
	WaralabaName *string `query:"waralaba_name"`
	CategoryId   *string `query:"category_id"`
	Search       *string `query:"search"`
	Page         int     `query:"page"`
	PageSize     int     `query:"page_size"`
	OrderBy      string  `query:"order_by"`
	OrderType    string  `query:"order_type"`
}

type UpdateCategoryRequest struct {
	Id              int64     `json:"id" validate:"required"`
	CategoryId      int64     `json:"category_id" validate:"required"`
	WaralabaName    string    `json:"waralaba_name" validate:"required"`
	Prize           string    `json:"prize" validate:"required"`
	Contact         string    `json:"contact" validate:"required"`
	BrochureLink    string    `json:"brochure_link" validate:"required"`
	Since           time.Time `json:"since" validate:"required"`
	OutletTotal     int64     `json:"outlet_total" validate:"required"`
	LicenseDuration string    `json:"license_duration" validate:"required"`
}

func (m *UpdateCategoryRequest) Sanitize() {
	m.WaralabaName = sanitizer.SanitizeHTML(m.WaralabaName)
	m.Prize = sanitizer.SanitizeHTML(m.Prize)
	m.Contact = sanitizer.SanitizeHTML(m.Contact)
	m.BrochureLink = sanitizer.SanitizeHTML(m.BrochureLink)
}
