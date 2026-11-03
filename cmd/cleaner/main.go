package main

import (
	"fmt"
	"os"

	"github.com/dmitry-konchenko/go-data-cleaner.git/internal/cleaner"
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

	records, err := cleaner.ReadCSV(inputPath)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		os.Exit(1)
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		record[0] = cleaner.NormalizeName(record[0])
		record[1] = cleaner.NormalizeEmail(record[1])
		record[2] = cleaner.NormalizePhone(record[2])
	}
	duplicates := cleaner.FindDuplicates(records)

	if len(duplicates) > 0 {
		fmt.Println("\nDuplicates found:")

		for _, duplicate := range duplicates {
			fmt.Println("-", duplicate)
		}
	} else {
		fmt.Println("\nNo duplicates found.")
	}
	err = cleaner.WriteCSV(outputPath, records)
	if err != nil {
		fmt.Println("Error writing CSV:", err)
		os.Exit(1)
	}
	reportPath := "testdata/report.txt"

	err = cleaner.WriteReport(
		reportPath,
		len(records)-1,
		duplicates,
	)
	if err != nil {
		fmt.Println("Error writing report:", err)
		os.Exit(1)
	}

	fmt.Println("Report saved to:", reportPath)
	fmt.Println("Done!")
	fmt.Println("Saved to:", outputPath)
}
