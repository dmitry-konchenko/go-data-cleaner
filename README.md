# Go Data Cleaner

A small command-line utility written in Go for cleaning and validating CSV data.

The application reads a CSV file, normalizes user data, detects duplicates, skips invalid rows, and generates a cleaned CSV file together with a processing report.

## Features

- CSV file reading and writing
- Name normalization
- Email normalization
- Russian phone number normalization
- Duplicate detection by email and phone
- Invalid row handling
- Processing report generation
- Unit tests

## Example

Input:

```csv
name,email,phone
иванов иван иванович,IVAN@MAIL.RU,89991234567
ПЕТРОВ ПЕТР ПЕТРОВИЧ,petr@example.com,+7 999 888-77-66
Иванов Иван,ivan@mail.ru,+7 999 123-45-67
```

Output:

```csv
name,email,phone
Иванов Иван Иванович,ivan@mail.ru,+79991234567
Петров Петр Петрович,petr@example.com,+79998887766
Иванов Иван,ivan@mail.ru,+79991234567
```

The application also detects duplicate emails and phone numbers.

Example report:

```text
Data Cleaner Report

Total records: 3
Invalid rows: 0
Duplicates found: 2

- Duplicate email: ivan@mail.ru
- Duplicate phone: +79991234567
```

## Project structure

```text
go-data-cleaner/
├── cmd/
│   └── cleaner/
│       └── main.go
├── internal/
│   └── cleaner/
│       ├── cleaner.go
│       └── cleaner_test.go
├── testdata/
│   └── input.csv
├── .gitignore
├── README.md
└── go.mod
```

## Run

```bash
go run ./cmd/cleaner testdata/input.csv testdata/output.csv
```

The cleaned file will be saved to:

```text
testdata/output.csv
```

The processing report will be saved to:

```text
testdata/report.txt
```

## Tests

Run all tests:

```bash
go test ./...
```

Run tests with detailed output:

```bash
go test -v ./...
```

## Technologies

- Go
- Standard library
- encoding/csv
- Unit testing with Go `testing` package

## Future improvements

- Email validation
- More flexible phone number validation
- Configurable CSV columns
- JSON support
- Command-line flags
- Detailed error reporting