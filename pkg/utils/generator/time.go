package generator

import (
	"log"
	"time"
)

func GenerateTimeNowJakarta() time.Time {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
	}
	now := time.Now().In(loc)
	return now
}

func GenerateTimeNowLocal(loc *time.Location) time.Time {
	now := time.Now().In(loc)
	return now
}
