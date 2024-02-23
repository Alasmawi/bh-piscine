// package piscine

// import "fmt"

// func QuadA(x, y int) {
// 	w := x
// 	h := y
// 	for i := 0; i <= h-2; i++ {
// 		if i == 0 || i <= h-2 {
// 			fmt.Print("o")

// 		} else {
// 			fmt.Print("|")
// 		}
// 		for j := 0; j <= w-2; j++ {
// 			if i == 0 || i <= h-2 {

// 				fmt.Print("-")
// 			} else {
// 				fmt.Print(" ")
// 			}

// 		}
// 		if i == 0 || i <= h-2 {
// 			fmt.Print("o")

// 		} else {
// 			fmt.Print("|")
// 		}
// 	fmt.Println()}

// 	// Ending line after each row

// }
package piscine

import "fmt"

func QuadA(x, y int) {
	if x > 0 || y > 0 {

		for height := 0; height < y; height++ {

			for width := 0; width < x; width++ {

				if (height == 0 || height == y-1) && (width == 0 || width == x-1) {
					fmt.Print("o")
				} else if height == 0 || height == y-1 {
					fmt.Print("-")
				} else if width == 0 || width == x-1 {
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}
