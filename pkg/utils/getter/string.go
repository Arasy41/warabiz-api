package getter

func GetStringLength(str string) int {
	return len([]rune(str))
}
