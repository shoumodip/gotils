package main

import (
	"fmt"
	"os"
	"bufio"
)

func printFile(filePath string) {
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

	for i := len(lines); i > 0; i-- {
		fmt.Println(lines[i - 1])
	}
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		printFile(os.Args[i])
	}
}
