package piscine

func SplitWhiteSpaces(s string) []string {
	arr1 := []string{}
	s += " "
	temp := ""
	if len(s) == 0 {
		return arr1
	}
	for _, v := range s {
		if v == ' ' || v == '\t' || v == '\n' {
			if temp != "" {
				arr1 = append(arr1, temp)
				temp = ""
			}
		} else {
			temp += string(v)
		}
	}
	return arr1
}
