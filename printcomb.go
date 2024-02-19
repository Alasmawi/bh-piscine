package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for k := '0'; k < '9'; k++ {
		for j := '0'; j < '9'; j++ {
			for i := '0'; i <= '9'; i++ {
				if k < j && j < i {

					z01.PrintRune(k)
					z01.PrintRune(j)
					z01.PrintRune(i)
					if k == '7' && j == '8' && i == '9' {
						z01.PrintRune('\n')
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
			}
		}
	}
}
