package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	st := ""
	for _, str := range a {
		st = str
		for _, chr := range st {
			z01.PrintRune(chr)
		}
		z01.PrintRune('\n')
	}
}
