package piscine

func IsUpper(s string) bool {
	boolean := false
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 67 && rune(s[i]) <= 90 {
			boolean = true
		}
	}
	return boolean
}
