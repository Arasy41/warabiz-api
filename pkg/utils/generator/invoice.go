package generator

import (
	"fmt"
	"time"

	"warabiz/api/pkg/utils/converter"
)

func GenerateInvoice(userID int64, pattern string) string {
	timeNow := time.Now().UTC()
	month := converter.ConvertMonthtoRoman(int(timeNow.Month()))
	invoice := fmt.Sprintf("%s%v/%s/%v", pattern, userID, month, timeNow.Unix())
	return invoice
}
