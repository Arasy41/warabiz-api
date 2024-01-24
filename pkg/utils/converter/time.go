package converter

import (
	"time"

	constants "warabiz/api/pkg/constants/general"
)

var (
	DefaultLoc, _ = time.LoadLocation("Asia/Jakarta")
)

func GetTimeLocation(loc string) *time.Location {
	Loc, err := time.LoadLocation(loc)
	if err != nil {
		return DefaultLoc
	}
	return Loc
}

func ConvertTimeToTicksWindows(time time.Time) int64 {
	return ((time.Unix()) * 10000000) + 621355968000000000
}

func ConvertTicksWindowsToTime(ticks int64) time.Time {
	base := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	return time.Unix(ticks/10000000+base, ticks%10000000).UTC()
}

func ConvertToDateTime(dateTime string) (time.Time, error) {
	DateTime, err := time.ParseInLocation(constants.DBTimeLayout, dateTime, DefaultLoc)
	if err != nil {
		return time.Time{}, err
	}
	return DateTime, err
}

func ConvertToDate(dateTime string) (time.Time, error) {
	DateTime, err := time.ParseInLocation(constants.LayoutDateOnly, dateTime, DefaultLoc)
	if err != nil {
		return time.Time{}, err
	}
	return DateTime, err
}

func ConvertToDateTimeRequest(dateTime string) (time.Time, error) {
	DateTime, err := time.ParseInLocation(constants.RequestTimeLayout, dateTime, DefaultLoc)
	if err != nil {
		return time.Time{}, err
	}
	return DateTime, err
}

func ConvertTimeToJakartaTime(t time.Time) time.Time {
	return t.In(DefaultLoc)
}

func ConvertTimeToLocal(t time.Time, loc *time.Location) time.Time {
	return t.In(loc)
}
