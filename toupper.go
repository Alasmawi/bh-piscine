package piscine

func ToUpper(s string) string {
	arr := []rune(s)
	arrr := ""
	for i := 0; i <= len(s); i++ {
		if arr[i] >= 97 && arr[i] <= 122 {
			arr[i] = arr[i] - 32
		}
	}
	for b := 0; b < len(s); b++ {
		arrr = arrr + string(arr[b])
	}

	return arrr
}
