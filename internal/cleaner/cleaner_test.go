package cleaner

import "testing"

func TestNormalizeName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"иванов иван иванович", "Иванов Иван Иванович"},
		{"ПЕТРОВ ПЕТР ПЕТРОВИЧ", "Петров Петр Петрович"},
		{"  сидоров   сидор  сидорович  ", "Сидоров Сидор Сидорович"},
	}

	for _, test := range tests {
		result := NormalizeName(test.input)

		if result != test.expected {
			t.Errorf(
				"NormalizeName(%q) = %q; expected %q",
				test.input,
				result,
				test.expected,
			)
		}
	}
}

func TestNormalizeEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"IVAN@MAIL.RU", "ivan@mail.ru"},
		{" Petr@Example.COM ", "petr@example.com"},
		{"test@gmail.com", "test@gmail.com"},
	}

	for _, test := range tests {
		result := NormalizeEmail(test.input)

		if result != test.expected {
			t.Errorf(
				"NormalizeEmail(%q) = %q; expected %q",
				test.input,
				result,
				test.expected,
			)
		}
	}
}

func TestNormalizePhone(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"89991234567", "+79991234567"},
		{"+7 999 123-45-67", "+79991234567"},
		{"8 (999) 888-77-66", "+79998887766"},
		{"+79995554433", "+79995554433"},
	}

	for _, test := range tests {
		result := NormalizePhone(test.input)

		if result != test.expected {
			t.Errorf(
				"NormalizePhone(%q) = %q; expected %q",
				test.input,
				result,
				test.expected,
			)
		}
	}
}

func TestFindDuplicates(t *testing.T) {
	records := [][]string{
		{"name", "email", "phone"},
		{"Иванов Иван", "ivan@mail.ru", "+79991234567"},
		{"Петров Петр", "petr@mail.ru", "+79998887766"},
		{"Иванов И.", "ivan@mail.ru", "+79991234567"},
	}

	duplicates := FindDuplicates(records)

	expectedCount := 2

	if len(duplicates) != expectedCount {
		t.Errorf(
			"FindDuplicates returned %d duplicates; expected %d",
			len(duplicates),
			expectedCount,
		)
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"TEST@EXAMPLE.COM", true},
		{"invalid-email", false},
		{"test@", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsValidEmail(test.email)

		if result != test.expected {
			t.Errorf(
				"IsValidEmail(%q) = %v; expected %v",
				test.email,
				result,
				test.expected,
			)
		}
	}
}

func TestIsValidPhone(t *testing.T) {
	tests := []struct {
		phone    string
		expected bool
	}{
		{"89991234567", true},
		{"+7 999 123-45-67", true},
		{"8 (999) 888-77-66", true},
		{"+79995554433", true},

		{"12345", false},
		{"", false},
		{"abc", false},
		{"+7999123456", false},
		{"+799912345678", false},
	}

	for _, test := range tests {
		result := IsValidPhone(test.phone)

		if result != test.expected {
			t.Errorf(
				"IsValidPhone(%q) = %v; expected %v",
				test.phone,
				result,
				test.expected,
			)
		}
	}
}
