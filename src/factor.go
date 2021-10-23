package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintln(os.Stderr, "  factor NUMBER")
		fmt.Fprintln(os.Stderr, "Error: insufficient arguments")
		os.Exit(1)
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid number '%s'\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Printf("%d:", number)
	for number > 1 {
		for i := 2; i <= number; i++ {
			if number % i == 0 {
				fmt.Printf(" %d", i)
				number /= i
				break
			}
		}
	}
}
