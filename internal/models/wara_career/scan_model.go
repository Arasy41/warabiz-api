package wara_career

import (
	"context"
	"time"
	"warabiz/api/pkg/http/locals"
	"warabiz/api/pkg/utils/converter"
)

type WaraCareer struct {
	Id          int64      `json:"id"`
	CareerTitle string     `json:"career_title"`
	Description string     `json:"description"`
	Address     string     `json:"address"`
	ImageUrl    string     `json:"image_url"`
	CreatedBy   string     `json:"created_by"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	UpdatedBy   *string    `json:"updated_by"`
}

func (m *WaraCareer) SetTimeToLocal(ctx context.Context) {
	timeLoc := locals.GetTimeLoc(ctx)
	m.CreatedAt = converter.ConvertTimeToLocal(m.CreatedAt, converter.GetTimeLocation(timeLoc))
	if m.UpdatedAt != nil {
		*m.UpdatedAt = converter.ConvertTimeToLocal(*m.UpdatedAt, converter.GetTimeLocation(timeLoc))
	}
}

type WaraCareerList struct {
	Id          int64  `json:"id"`
	CareerTitle string `json:"career_title"`
	Description string `json:"description"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
}
