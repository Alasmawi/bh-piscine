package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// save at index 0
	name := os.Args[0] // array of strings
	// print at index from index
	for _, str := range name[2:] {
		z01.PrintRune(str)
		z01.PrintRune('\n')
	}
}
