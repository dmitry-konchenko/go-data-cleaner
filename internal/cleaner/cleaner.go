package cleaner

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func ReadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Разрешаем строкам содержать разное количество полей.
	// Корректность строк проверяем самостоятельно.
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func NormalizeName(name string) string {
	words := strings.Fields(strings.ToLower(name))

	for i, word := range words {
		runes := []rune(word)

		if len(runes) > 0 {
			runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]
		}

		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}

func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func NormalizePhone(phone string) string {
	var digits strings.Builder

	for _, r := range phone {
		if r >= '0' && r <= '9' {
			digits.WriteRune(r)
		}
	}

	result := digits.String()

	if len(result) == 11 && result[0] == '8' {
		result = "7" + result[1:]
	}

	if len(result) == 11 && result[0] == '7' {
		return "+" + result
	}

	return result
}

func FindDuplicates(records [][]string) []string {
	seenEmails := make(map[string]bool)
	seenPhones := make(map[string]bool)

	var duplicates []string

	for i, record := range records {
		if i == 0 {
			continue
		}

		if len(record) < 3 {
			continue
		}

		email := record[1]
		phone := record[2]

		if seenEmails[email] {
			duplicates = append(
				duplicates,
				"Duplicate email: "+email,
			)
		} else {
			seenEmails[email] = true
		}

		if seenPhones[phone] {
			duplicates = append(
				duplicates,
				"Duplicate phone: "+phone,
			)
		} else {
			seenPhones[phone] = true
		}
	}

	return duplicates
}

func WriteCSV(filePath string, records [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(records)
}

func WriteReport(
	filePath string,
	totalRecords int,
	invalidRows int,
	duplicates []string,
) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(
		file,
		"Data Cleaner Report\n\n"+
			"Total records: %d\n"+
			"Invalid rows: %d\n"+
			"Duplicates found: %d\n\n",
		totalRecords,
		invalidRows,
		len(duplicates),
	)

	if err != nil {
		return err
	}

	if len(duplicates) == 0 {
		_, err = fmt.Fprintln(file, "No duplicates found.")
		return err
	}

	for _, duplicate := range duplicates {
		if _, err := fmt.Fprintln(file, "-", duplicate); err != nil {
			return err
		}
	}

	return nil
}
