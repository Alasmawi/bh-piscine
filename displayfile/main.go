package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("quest8.txt")
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Print(string(dat[2:]))
}
