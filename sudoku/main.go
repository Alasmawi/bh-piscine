package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	//solve when valid
	if vaild() == true {
		array := [9][9]rune{}
		for i := 1; i <= 9; i++ {
			x := 0
			for _, c := range os.Args[i] {
				array[i-1][x] = c
				x++
			}
		} //checks for dudps in grid
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if array[i][j] >= '1' && array[i][j] <= '9' {
					if Isdubl(array, array[i][j], i, j) == false {
						fmt.Println("Error")
						return
					}
				}
			}
		}
		//check if solvable
		if rslv(array) == true {
			//solves the grid
			array = solve(array)
			array = solve(array)
			//print grid
			print(array)
		} else {
			fmt.Println("Error")
		}
	}
}

// checks conditions that return errors
func vaild() bool {
	//verify no. of col =9
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return false
	} else {
		//verify no of rows
		for i := 1; i <= 9; i++ {
			if len(os.Args[i]) != 9 {
				fmt.Println("Error")
				return false
			}
		}
		//verify grid is filled
		for i := 1; i <= 9; i++ {
			for _, z := range os.Args[i] {
				if z != '.' {
					if !(z >= '1' && z <= '9') {
						fmt.Println("Error")
						return false
					}
				}
			}
		}
	}
	//return true when no errors are
	return true
}

// prints grid
func print(array [9][9]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune((array[i][j]))
			if j < 8 {
				z01.PrintRune(' ')
			}
		}
		fmt.Print("\n")
	}
}

// checks if givin num is in row
func inRow(array [9][9]rune, num rune, row int) bool {
	for i := 0; i < 9; i++ {
		if array[row][i] == num {
			return false
		}
	}
	return true
}

// checks if givin num is in col
func inCol(array [9][9]rune, num rune, col int) bool {
	for i := 0; i < 9; i++ {
		if array[i][col] == num {
			return false
		}
	}
	return true
}

// checks if givin num is in box
func inBox(array [9][9]rune, num rune, col int, row int) bool {
	row = row - row%3
	col = col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if array[row+i][col+j] == num {
				return false
			}
		}
	}
	return true
}

// checks if a number can be placed in a certain cell
func isTrue(array [9][9]rune, num rune, row int, col int) bool {
	if inBox(array, num, col, row) && inCol(array, num, col) && inRow(array, num, row) {
		return true
	}
	return false
}

// backtracking
func rslv(array [9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if array[i][j] == '.' {
				for v := '1'; v <= '9'; v++ {
					if isTrue(array, v, i, j) == true {
						array[i][j] = v
						if rslv(array) == true {
							return true
						} else {
							array[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
} //fills the the array with correct values according to sudoku rules
func solve(array [9][9]rune) [9][9]rune {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if array[i][j] == '.' {
				for v := '1'; v <= '9'; v++ {
					if isTrue(array, v, i, j) == true {
						array[i][j] = v
						if rslv(array) != true {
							array[i][j] = '.'
						}
					}
				}
			}
		}
	}
	return array
}

// checks if a number is duplicated in its row, column, or box
func Isdubl(array [9][9]rune, num rune, row int, col int) bool {
	for i := 0; i < 9; i++ {
		if i != col {
			if array[row][i] == num {
				return false
			}
		}
	}
	for i := 0; i < 9; i++ {
		if i != row {
			if array[i][col] == num {
				return false
			}
		}
	}
	x := row - row%3
	z := col - col%3
	for i := x; i < 3; i++ {
		for j := z; j < 3; j++ {
			if row != i && col != j {
				if array[i][j] == num {
					return false
				}
			}
		}
	}
	return true
}
