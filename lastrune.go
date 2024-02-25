package piscine

func LastRune(s string) rune {
	arr := []rune(s)

	return arr[len(s)-1]
}
