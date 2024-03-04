package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 1
	for _, str := range tab {
		if f(str) {
			count++
		}
	}
	if count == 1 {
		return 0
	} else {
		return count
	}
}
