package piscine

func ToUpper(s string) string {
	arr := []rune(s)
	arrr := ""
	for i := 0; i <= len(s); i++ {
		if arr[i] >= 97 && arr[i] <= 122 {
			arr[i] = arr[i] - 32
		}

		arrr = string(arr)
	}

	return arrr
}
