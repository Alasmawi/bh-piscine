package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args
	for _, str := range name {
		for _, ch := range str {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}
