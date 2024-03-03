package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Print(string(dat[2:]))
}
