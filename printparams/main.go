package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args
	for _, str := range name { // enter each string
		for _, ch := range str { // print each rune
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
