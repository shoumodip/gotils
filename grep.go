package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
)

func matchFile(filePath string, pattern *regexp.Regexp) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not read file '%s': %s\n", os.Args[1], err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	lineNr := 1

	for scanner.Scan() {
		line := scanner.Text()
		match := pattern.FindStringIndex(line)

		if len(match) == 2 {
			fmt.Printf("%s:%d:%d: %s\n", filePath, lineNr, match[0] + 1, line)
		}

		lineNr++
	}

	file.Close()
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintln(os.Stderr, "  grep PATTERN [FILE]")
		fmt.Fprintln(os.Stderr, "Error: insufficient arguments")
		os.Exit(1)
	}

	pattern, err := regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid pattern '%s': %s\n", os.Args[1], err)
		os.Exit(1)
	}

	for i := 2; i < len(os.Args); i++ {
		filePath := os.Args[i]
		matchFile(filePath, pattern)
	}
}
