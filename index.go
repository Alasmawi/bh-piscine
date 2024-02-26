package piscine

func Index(s string, toFind string) int {
	count := 1
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == rune(toFind[i]) {
			x := 0
			for j := 0; j < len(toFind); j++ {

				if rune(s[x]) == rune(toFind[j]) {
					count++
					x++
				}
				if count+1 == len(toFind) {
					count = i
				}
			}
		}
	}
	if count != 0 {
		return count
	} else {
		return -1
	}
}
