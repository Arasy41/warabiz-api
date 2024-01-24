package warabiz

import (
	"context"
	"time"
	"warabiz/api/pkg/http/locals"
	"warabiz/api/pkg/utils/converter"
)

type Warabiz struct {
	Id              int64      `json:"id"`
	CategoryId      int64      `json:"category_id"`
	CategoryName    string     `json:"category_name"`
	WaralabaName    string     `json:"waralaba_name"`
	Prize           string     `json:"prize"`
	Contact         string     `json:"contact"`
	BrochureLink    string     `json:"brochure_link"`
	Since           time.Time  `json:"since"`
	OutletTotal     int64      `json:"outlet_total"`
	LicenseDuration string     `json:"license_duration"`
	CreatedBy       string     `json:"created_by"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	UpdatedBy       *string    `json:"updated_by"`
}

func (m *Warabiz) SetTimeToLocal(ctx context.Context) {
	timeLoc := locals.GetTimeLoc(ctx)
	m.CreatedAt = converter.ConvertTimeToLocal(m.CreatedAt, converter.GetTimeLocation(timeLoc))
	if m.UpdatedAt != nil {
		*m.UpdatedAt = converter.ConvertTimeToLocal(*m.UpdatedAt, converter.GetTimeLocation(timeLoc))
	}
}

type WarabizList struct {
	Id           int64  `json:"id"`
	CategoryId   int64  `json:"category_id"`
	WaralabaName string `json:"waralaba_name"`
	Prize        string `json:"prize"`
}
