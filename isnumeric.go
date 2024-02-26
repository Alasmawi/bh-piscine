package piscine

func IsNumeric(s string) bool {
	boolean := false
	if rune(s[0]) == ' ' {
		return false
	}
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
			boolean = true
		} else {
			boolean = false
			break
		}
	}
	return boolean
}
