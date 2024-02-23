package piscine

import "fmt"

func QuadC(x, y int) {
	if x > 0 || y > 0 {

		for height := 0; height < y; height++ {

			for width := 0; width < x; width++ {

				if (height == 0 || height == y-1) && (width == 0 || width == x-1) {
					if height == 0 && (width == 0 || width == x-1) {
						fmt.Print("A")
					} else if height == y-1 && (width == 0 || width == x-1) {
						fmt.Print("C")
					}
				} else if height == 0 || height == y-1 {
					fmt.Print("B")
				} else if width == 0 || width == x-1 {
					fmt.Print("B")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}
