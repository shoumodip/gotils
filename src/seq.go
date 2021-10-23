package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseInt(s string, label string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid %s limit '%s'\n", label, s)
		os.Exit(1)
	}
	return i
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  seq [FIRST] LAST")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		fmt.Fprintln(os.Stderr, "Error: insufficient arguments")
		os.Exit(1)
	}

	first, last := 0, 0

	if len(os.Args) == 2 {
		last = parseInt(os.Args[1], "last")
	} else if len(os.Args) == 3 {
		first = parseInt(os.Args[1], "first")
		last = parseInt(os.Args[2], "last")
	} else {
		usage()
		fmt.Fprintln(os.Stderr, "Error: invalid number of arguments:", len(os.Args) - 1)
		os.Exit(1)
	}

	for i := first; i <= last; i++ {
		fmt.Println(i)
	}
}
