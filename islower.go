package piscine

func IsLower(s string) bool {
	boolean := false
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			boolean = true
		} else {
			boolean = false
			break
		}
	}
	return boolean
}
