package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args

	for i := 1; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	for _, str := range arr { // enter each string
		for _, ch := range str { // print each rune
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
