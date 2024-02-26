package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	s := string(n)
	arr := []rune(s)

	for i := 0; i < len(arr); i++ {
		max := i
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[j] > arr[max] {
				// changing the index to show the min value
				max = j
			}
		}
		temp := arr[i]
		arr[i] = arr[max]
		arr[max] = temp

		for i := 0; i < len(arr); i++ {
			z01.PrintRune(arr[i])
		}
	}
}
