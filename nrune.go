package piscine

func NRune(s string, n int) rune {
	max := len(s)
	if n < 0 {
		return 0
	} else if n > max {
		return 0
	} else if n == 0 {
		return rune(s[0])
	} else {
		return rune(s[n-1])
	}
}
