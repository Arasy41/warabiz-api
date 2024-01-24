package generator

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenerateRandomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func GenerateRandomBool() bool {
	return rand.Intn(2) == 1
}

func GenerateRandomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func GenerateRandomCharWithLenght(length int, charset string) string {
	chars := []rune(charset)
	rand.Seed(time.Now().UnixNano())
	code := make([]rune, length)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}

func GenerateRandomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func GenerateRandomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func GenerateUUID() string {
	return uuid.New().String()
}
