package piscine

func AlphaCount(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if (rune(s[i]) >= 67 && rune(s[i]) <= 90) || (rune(s[i]) >= 97 && rune(s[i]) <= 122) {
			count++
		}
	}
	if count == 0 {
		return 0
	} else {
		return count + 1
	}
}
