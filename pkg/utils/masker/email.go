package masker

import (
	"fmt"
	"strings"
)

func GenerateMaskEmail(email string) string {
	if email == "" {
		return ""
	}
	array := strings.Split(email, "@")
	str1 := array[0]
	str2 := array[1]
	length := len(str1)
	strs := str1[0:1]
	stre := str1[length-1 : length]
	var mask string
	for i := 0; i < length-2; i++ {
		mask = mask + "*"
	}
	return fmt.Sprintf("%s%s%s@%s", strs, mask, stre, str2)
}
