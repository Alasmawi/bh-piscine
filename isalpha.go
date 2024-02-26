package piscine

func IsAlpha(s string) bool {
	boolean := false
	if rune(s[0]) == ' ' {
		return false
	}
	for i := 0; i < len(s); i++ {
		if ((rune(s[i]) >= 48 && rune(s[i]) <= 57) || rune(s[i]) >= 65 && rune(s[i]) <= 90) || (rune(s[i]) >= 97 && rune(s[i]) <= 122) {
			boolean = true
		} else {
			boolean = false
			break
		}
	}
	return boolean
}
