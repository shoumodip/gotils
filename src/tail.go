package main

import (
	"fmt"
	"os"
	"bufio"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		filePath := os.Args[i]
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: could not read file '%s': %s\n", filePath, err)
			os.Exit(1)
		}

		scanner := bufio.NewScanner(file)
		lines := []string{}

		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		file.Close()

		for i := max(len(lines) - 10, 0); i < len(lines); i++ {
			fmt.Println(lines[i])
		}
	}
}
