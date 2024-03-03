package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("File name missing")
		os.Exit(0)
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(0)
	}
	Fname := os.Args[1]
	dat, err := os.ReadFile(Fname)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Print(string(dat[2:]))
}
