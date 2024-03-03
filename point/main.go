package main

import "github.com/01-edu/z01"

func setPoint(pointer *point) {
	pointer.x = 42
	pointer.y = 21
}

type point struct {
	x, y int
}

func main() {
	points := &point{}
	x := ' '
	for i := 1; i <= 14; i++ {
		switch i {
		case 1:
			x = 'x'
		case 2:
			x = ' '
		case 3:
			x = '='
		case 4:
			x = ' '
		case 5:
			x = 52
		case 6:
			x = 50
		case 7:
			x = ','
		case 8:
			x = ' '
		case 9:
			x = 'y'
		case 10:
			x = ' '
		case 11:
			x = '='
		case 12:
			x = ' '
		case 13:
			x = 50
		case 14:
			x = 49

		}
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
	setPoint(points)

}
