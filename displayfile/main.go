package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("File name missing")
		return
	}
	fmt.Print(string(dat[2:]))
}
