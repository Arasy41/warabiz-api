package getter

import (
	"fmt"
	"math"
)

func GetHundredsCategory(number float64) string {
	min := math.Floor(float64(number)/100) * 100
	return fmt.Sprintf("%v-%v", min+1, min+100)
}

func GetThousandsCategory(number float64) string {
	min := math.Floor(float64(number)/1000) * 1000
	return fmt.Sprintf("%v-%v", min+1, min+1000)
}
