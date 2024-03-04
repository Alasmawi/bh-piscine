package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	num := 0
	for i := 0; i < len(nbr); i++ {
		num = num*len(baseFrom) + indexof(baseFrom, nbr[i])
	}
	result := ""
	for num > 0 {
		remainder := num % len(baseTo)
		result = string(baseTo[remainder]) + result
		num = num / len(baseTo)
	}
	return result
}

func indexof(str string, char byte) int {
	for i := 0; i < len(str); i++ {
		if str[i] == char {
			return i
		}
	}
	return -1
}
