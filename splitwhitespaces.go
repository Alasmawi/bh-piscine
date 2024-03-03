package piscine

func SplitWhiteSpaces(s string) []string {
	arr1 := []string{}
	index := 0
	if len(s) == 0 {
		return arr1
	}
	for v := range s {
		if v == ' ' || v == '\t' || v == '\n' {
			index++
		} else {
			arr1[index] += string(v)
		}
	}
	return arr1
}
