package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		dat, err := os.ReadFile(os.Args[4])
		if err != nil {
			fmt.Println("File name missing")
			return
		}
		fmt.Print(string(dat[2:]))
	} else {
		fmt.Println("File name missing")
	}
}
