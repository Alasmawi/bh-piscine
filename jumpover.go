package piscine

func JumpOver(str string) string {
	if len(str) <= 2 {
		return "\n"
	}
	arr := []rune(str)
	str1 := "\n"
	for i := 2; i < len(str); i += 3 {
		str1 += string(arr[i])
	}
	return str1
}
