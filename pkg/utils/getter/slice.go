package getter

func GetCleanSliceOfString(strings []string) []string {
	seen := make(map[string]struct{}, len(strings))
	j := 0
	for _, v := range strings {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			strings[j] = v
			j++
		}
	}
	return strings[:j]
}
