package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	arr := []rune(string(n))

	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		z01.PrintRune(arr[i])
	}
}
