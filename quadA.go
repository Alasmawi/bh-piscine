package piscine

import "fmt"

func QuadA(x, y int) {
	for i := 0; i <= x-2; i++ {
		if i == 0 || i == x-2 {
			fmt.Print("o")
		} else {
			fmt.Print("|")
		}
		for j := 0; j <= y-1; j++ {
			// print
			if i == 0 || i == x-2 {
				fmt.Print("-")
			} else {
				fmt.Print(" ")
			}
		}
		if i == 0 || i == x-2 {
			fmt.Print("o")
		}
		if i <= y {
			fmt.Print("|")
		}
		// Ending line after each row
		fmt.Println()
	}
}
