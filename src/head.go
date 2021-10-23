package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		filePath := os.Args[i]
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: could not read file '%s': %s\n", filePath, err)
			os.Exit(1)
		}

		scanner := bufio.NewScanner(file)
		head := []string{}

		for i := 0; scanner.Scan() && i < 10; i++ {
			head = append(head, scanner.Text())
		}

		file.Close()

		for _, line := range head {
			fmt.Println(line)
		}
	}
}
