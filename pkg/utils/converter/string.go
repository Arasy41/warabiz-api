package converter

import (
	"regexp"
	"strings"
	"unicode"
)

func removeDuplicateUnderscores(s string) string {
	var result strings.Builder
	underscoreCount := 0

	for _, char := range s {
		if char == '_' {
			underscoreCount++
			if underscoreCount > 1 {
				continue
			}
		} else {
			underscoreCount = 0
		}

		result.WriteRune(char)
	}

	return result.String()
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ConvertToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	str = strings.ToLower(snake)
	var result strings.Builder
	for i, char := range str {
		if unicode.IsDigit(char) {
			if i > 0 && !unicode.IsDigit(rune(str[i-1])) {
				result.WriteRune('_')
			}
			result.WriteRune(char)
			if i < len(str)-1 && !unicode.IsDigit(rune(str[i+1])) {
				result.WriteRune('_')
			}
		} else {
			result.WriteRune(char)
		}
	}
	return removeDuplicateUnderscores(result.String())
}
