package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go-data-cleaner <input.csv> <output.csv>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	fmt.Println("Input:", inputPath)
	fmt.Println("Output:", outputPath)
}
