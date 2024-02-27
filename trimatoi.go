package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	for _, r := range s {
		if r == '-' && result == 0 {
			sign = -1
		}
		if r-'0' >= 0 && r-'0' <= 9 {
			digit := int(r - '0')

			result = result*10 + digit
		}
	}
	return result * sign
}
