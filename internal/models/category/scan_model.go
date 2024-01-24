package category

import (
	"context"
	"time"
	"warabiz/api/pkg/http/locals"
	"warabiz/api/pkg/utils/converter"
)

type Category struct {
	Id           int64      `json:"id"`
	CategoryName int64      `json:"category_name"`
	CreatedBy    string     `json:"created_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	UpdatedBy    *string    `json:"updated_by"`
}

func (m *Category) SetTimeToLocal(ctx context.Context) {
	timeLoc := locals.GetTimeLoc(ctx)
	m.CreatedAt = converter.ConvertTimeToLocal(m.CreatedAt, converter.GetTimeLocation(timeLoc))
	if m.UpdatedAt != nil {
		*m.UpdatedAt = converter.ConvertTimeToLocal(*m.UpdatedAt, converter.GetTimeLocation(timeLoc))
	}
}

type CategoryList struct {
	Id           int64  `json:"id"`
	CategoryName string `json:"category_name"`
}
