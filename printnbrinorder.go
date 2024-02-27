package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	arr := []rune{}

	for n > 0 {
		temp := n % 10
		arr = append(arr, rune(temp+'0'))
		n = n / 10
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		z01.PrintRune(arr[i])
	}
}
