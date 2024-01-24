package converter

import (
	"strings"
	"time"
)

var hariIndonesia = map[string]string{
	"Sunday":    "Minggu",
	"Monday":    "Senin",
	"Tuesday":   "Selasa",
	"Wednesday": "Rabu",
	"Thursday":  "Kamis",
	"Friday":    "Jumat",
	"Saturday":  "Sabtu",
}

func ConvertTimeToDate1(t time.Time) string {
	var formattedDate string
	formattedDate = t.Format("Monday, 2 January 2006")
	for enDay, idDay := range hariIndonesia {
		formattedDate = strings.ReplaceAll(formattedDate, enDay, idDay)
	}
	return formattedDate
}
