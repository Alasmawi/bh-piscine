package piscine

func IsPrintable(s string) bool {
	boolean := false
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 32 && rune(s[i]) <= 125 {
			boolean = true
		} else {
			boolean = false
			break
		}
	}
	return boolean
}
