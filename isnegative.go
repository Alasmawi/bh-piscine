package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		print("T")
	} else {
		print("F")
	}

	z01.PrintRune('\n')
}
