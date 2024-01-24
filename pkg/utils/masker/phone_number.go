package masker

import (
	"strings"
)

func GenerateMaskPhoneNumber(phoneNumber string) string {
	if phoneNumber == "" {
		return ""
	}
	length := len(phoneNumber)
	str := phoneNumber[0:3]
	var str1 string
	if strings.Contains(str, "0") {
		str1 = phoneNumber[0:2]
	} else if strings.Contains(str, "+") {
		str1 = phoneNumber[0:4]
	} else {
		str1 = phoneNumber[0:3]
	}
	str2 := phoneNumber[length-2 : length]
	var str3 string
	for i := 0; i < length; i++ {
		if i == 0 || i == 1 {
			str3 = str3 + "a"
		} else if i == length-2 || i == length-1 {
			str3 = str3 + "b"
		} else {
			str3 = str3 + "*"
		}
	}
	str4 := strings.Replace(str3, "aa", str1, 1)
	str4 = strings.Replace(str4, "bb", str2, 1)
	return str4
}
