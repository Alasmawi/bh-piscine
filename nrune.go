package piscine

func NRune(s string, n int) rune {
	arr := []rune(s)
	if n < 1 || n > len(s) {
		return 0
	}
	return arr[n-1]
}
