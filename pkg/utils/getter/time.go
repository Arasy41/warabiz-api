package getter

import (
	"context"
	"warabiz/api/pkg/http/locals"
	"warabiz/api/pkg/utils/converter"
	"time"
)

func GetUTCStartDateEndDate(ctx context.Context, startedDateStr, endDateStr string) (*time.Time, *time.Time, error) {

	timeLoc := converter.GetTimeLocation(locals.GetTimeLoc(ctx))

	//* Parse Start date & set to UTC
	startDateLocal, err := time.ParseInLocation(time.RFC3339, startedDateStr, timeLoc)
	if err != nil {
		return nil, nil, err
	}
	startDateUTC := startDateLocal.In(time.UTC)

	//* Parse End date & set to UTC
	endDateLocal, err := time.ParseInLocation(time.RFC3339, endDateStr, timeLoc)
	if err != nil {
		return nil, nil, err
	}
	endDateUtc := endDateLocal.In(time.UTC)

	return &startDateUTC, &endDateUtc, nil
}
