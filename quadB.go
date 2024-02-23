package piscine

import "fmt"

func QuadB(x, y int) {
	if x > 0 || y > 0 {

		for height := 0; height < y; height++ {

			for width := 0; width < x; width++ {

				if (height == 0 || height == y-1) && (width == 0 || width == x-1) {
					if height == 0 && width == 0 {
						fmt.Print("/")
					} else if height == y-1 && width == x-1 {
						fmt.Print("/")

					} else {
						fmt.Print("\\")
					}

				} else if height == 0 || height == y-1 {
					fmt.Print("*")
				} else if width == 0 || width == x-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}
