package piscine

func Rot14(s string) string {
	//m+14=a  m-a
	//n+14=b
	//o+14=c
	//(letter - 'a')+14=rot 14 of a letter
	arr := []rune(s)
	str := ""
	for i := 0; i < len(s); i++ {
		if (arr[i] >= 'A' && arr[i] <= 'Z') || (arr[i] >= 'a' && arr[i] <= 'z') {
			if arr[i] >= 'A' && arr[i] <= 'Z' {
				if arr[i] >= 'M' {
					arr[i] = arr[i] - 12
				} else {
					arr[i] += 14
				}
			} else if arr[i] >= 'a' && arr[i] <= 'z' {
				if arr[i] > 'm' {
					arr[i] = arr[i] - 12
				} else {
					arr[i] += 14
				}
			}
		}

		str += string(arr[i])
	}
	return str
}
