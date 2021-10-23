package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintln(os.Stderr, "  yes TEXT")
		fmt.Fprintln(os.Stderr, "Error: insufficient arguments")
		os.Exit(1)
	}

	for {
		fmt.Println(os.Args[1])
	}
}
