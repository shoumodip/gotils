package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print(os.Args[i])
	}
	fmt.Println()
}
