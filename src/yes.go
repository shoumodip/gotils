package main

import (
	"fmt"
	"os"
)

func main() {
	text := ""

	if len(os.Args) == 1 {
		text = "y"
	} else {
		for i := 1; i < len(os.Args); i++ {
			if i > 1 {
				text += " "
			}
			text += os.Args[i]
		}
	}

	for {
		fmt.Println(text)
	}
}
