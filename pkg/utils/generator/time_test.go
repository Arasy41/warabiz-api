package generator

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	timeNow := time.Now()
	timeAfter := timeNow.Add(time.Minute * 5)
	subTime := timeAfter.Sub(timeNow)
	roundTime := math.Round(subTime.Minutes())
	fmt.Println(roundTime)
}
