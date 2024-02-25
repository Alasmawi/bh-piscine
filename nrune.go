package piscine

func NRune(s string, n int) rune {
	arr := []rune(s)
	max := len(s)
	if n < 0 {
		return 0
	} else if n > max {
		return 0
	} else if n != 0 {
		return arr[n-1]
	} else {
		return arr[0]
	}
}
