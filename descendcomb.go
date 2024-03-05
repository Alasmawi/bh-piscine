package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for k := '9'; k >= '0'; k-- {
		for j := '9'; j >= '0'; j-- {
			for i := '9'; i >= '0'; i-- {
				for o := '9'; o >= '0'; o-- {
					if !(k == '0' && j == '1' && i == '0' && o == '0') {
						z01.PrintRune(k)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(i)
						z01.PrintRune(o)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('\n')
					}
				}
			}

		}
	}
}
