package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printFile(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not read file '%s': %s\n", filePath, err)
		os.Exit(1)
	}

	fmt.Println(string(file))
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		printFile(os.Args[i])
	}
}
