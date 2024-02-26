package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == toFind {
			return i
		}
	}
	return -1
}
