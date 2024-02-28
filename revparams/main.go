package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1:]
	for i := len(name); i >= 0; i-- { // enter each string
		for _, ch := range name[i] { // print each rune
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
