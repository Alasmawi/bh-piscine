package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	s := string(n)
	arr := []rune(s)
	min := 0
	temp := ' '
	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				// changing the index to show the min value
				min = j
			}
		}
		temp = arr[i]
		arr[i] = arr[min]
		arr[min] = temp

		for i := 0; i < len(arr); i++ {
			z01.PrintRune(arr[i])
		}
	}
}
