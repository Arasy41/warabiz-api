package converter

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func ConvertFloatToRupiah(price float64) string {
	p := message.NewPrinter(language.Indonesian)
	moneyString := p.Sprintf("Rp %.2f", price)
	return moneyString
}
