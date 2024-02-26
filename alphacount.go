package piscine

func AlphaCount(s string) int {
	count := 1
	for i := 0; i < len(s); i++ {
		if (rune(s[i]) >= 67 && rune(s[i]) <= 90) || (rune(s[i]) >= 97 && rune(s[i]) <= 122) {
			count++
		}
	}
	return count
}
