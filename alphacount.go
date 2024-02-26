package piscine

func AlphaCount(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if (rune(s[i]) > 66 && rune(s[i]) < 91) || (rune(s[i]) > 96 && rune(s[i]) < 123) {
			count++
		}
	}
	return count
}
