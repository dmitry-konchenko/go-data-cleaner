package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dmitry-konchenko/go-data-cleaner.git/internal/cleaner"
)

func main() {
	inputPath := flag.String(
		"input",
		"",
		"path to input CSV file",
	)

	outputPath := flag.String(
		"output",
		"output.csv",
		"path to cleaned CSV file",
	)

	reportPath := flag.String(
		"report",
		"report.txt",
		"path to report file",
	)

	flag.Parse()

	if *inputPath == "" {
		fmt.Println("input file is required")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Input:", *inputPath)
	fmt.Println("Output:", *outputPath)
	fmt.Println("Report:", *reportPath)

	records, err := cleaner.ReadCSV(*inputPath)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		os.Exit(1)
	}

	if len(records) == 0 {
		fmt.Println("CSV file is empty")
		os.Exit(1)
	}

	invalidRows := 0

	cleanedRecords := [][]string{
		records[0],
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		if len(record) < 3 {
			fmt.Println("Invalid row:", record)
			invalidRows++
			continue
		}

		record[0] = cleaner.NormalizeName(record[0])
		record[1] = cleaner.NormalizeEmail(record[1])
		record[2] = cleaner.NormalizePhone(record[2])

		if !cleaner.IsValidEmail(record[1]) {
			fmt.Println("Invalid email:", record[1])
			invalidRows++
			continue
		}

		if !cleaner.IsValidPhone(record[2]) {
			fmt.Println("Invalid phone:", record[2])
			invalidRows++
			continue
		}

		cleanedRecords = append(cleanedRecords, record)
	}

	duplicates := cleaner.FindDuplicates(cleanedRecords)

	if len(duplicates) > 0 {
		fmt.Println("\nDuplicates found:")

		for _, duplicate := range duplicates {
			fmt.Println("-", duplicate)
		}
	} else {
		fmt.Println("\nNo duplicates found.")
	}

	err = cleaner.WriteCSV(*outputPath, cleanedRecords)
	if err != nil {
		fmt.Println("Error writing CSV:", err)
		os.Exit(1)
	}

	err = cleaner.WriteReport(
		*reportPath,
		len(records)-1,
		invalidRows,
		duplicates,
	)
	if err != nil {
		fmt.Println("Error writing report:", err)
		os.Exit(1)
	}

	fmt.Println("\nDone!")
	fmt.Println("Saved to:", *outputPath)
	fmt.Println("Report saved to:", *reportPath)
}
