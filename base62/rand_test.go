package xbase62

import (
	"strings"
	"testing"
)

func TestRand(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{name: "length 0", length: 0},
		{name: "length 1", length: 1},
		{name: "length 10", length: 10},
		{name: "length 50", length: 50},
		{name: "length 100", length: 100},
		{name: "negative length", length: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Rand(tt.length)

			// Check length
			if tt.length <= 0 {
				if result != "" {
					t.Errorf("Rand(%d) expected empty string, got %q", tt.length, result)
				}
				return
			}

			if len(result) != tt.length {
				t.Errorf("Rand(%d) length = %d; want %d", tt.length, len(result), tt.length)
			}

			// Check all characters are from Alphanumeric charset
			for _, char := range result {
				if !strings.ContainsRune(Alphanumeric, char) {
					t.Errorf("Rand(%d) contains invalid character %q", tt.length, char)
				}
			}
		})
	}
}

func TestRandNumeric(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{name: "length 0", length: 0},
		{name: "length 1", length: 1},
		{name: "length 10", length: 10},
		{name: "length 50", length: 50},
		{name: "negative length", length: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandNumeric(tt.length)

			// Check length
			if tt.length <= 0 {
				if result != "" {
					t.Errorf("RandNumeric(%d) expected empty string, got %q", tt.length, result)
				}
				return
			}

			if len(result) != tt.length {
				t.Errorf("RandNumeric(%d) length = %d; want %d", tt.length, len(result), tt.length)
			}

			// Check all characters are numeric
			for _, char := range result {
				if !strings.ContainsRune(Numeric, char) {
					t.Errorf("RandNumeric(%d) contains invalid character %q", tt.length, char)
				}
			}
		})
	}
}

func TestRandAlphabetic(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{name: "length 0", length: 0},
		{name: "length 1", length: 1},
		{name: "length 10", length: 10},
		{name: "length 50", length: 50},
		{name: "negative length", length: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandAlphabetic(tt.length)

			// Check length
			if tt.length <= 0 {
				if result != "" {
					t.Errorf("RandAlphabetic(%d) expected empty string, got %q", tt.length, result)
				}
				return
			}

			if len(result) != tt.length {
				t.Errorf("RandAlphabetic(%d) length = %d; want %d", tt.length, len(result), tt.length)
			}

			// Check all characters are alphabetic
			for _, char := range result {
				if !strings.ContainsRune(Alphabetic, char) {
					t.Errorf("RandAlphabetic(%d) contains invalid character %q", tt.length, char)
				}
			}
		})
	}
}

func TestRandUniqueness(t *testing.T) {
	const iterations = 100
	const length = 20

	t.Run("Rand uniqueness", func(t *testing.T) {
		results := make(map[string]bool)
		for range iterations {
			result := Rand(length)
			if results[result] {
				t.Errorf("Rand generated duplicate string: %q", result)
			}
			results[result] = true
		}
	})

	t.Run("RandNumeric uniqueness", func(t *testing.T) {
		results := make(map[string]bool)
		for range iterations {
			result := RandNumeric(length)
			if results[result] {
				t.Errorf("RandNumeric generated duplicate string: %q", result)
			}
			results[result] = true
		}
	})

	t.Run("RandAlphabetic uniqueness", func(t *testing.T) {
		results := make(map[string]bool)
		for range iterations {
			result := RandAlphabetic(length)
			if results[result] {
				t.Errorf("RandAlphabetic generated duplicate string: %q", result)
			}
			results[result] = true
		}
	})
}

func TestRandCharacterDistribution(t *testing.T) {
	const length = 1000

	t.Run("Rand uses full charset", func(t *testing.T) {
		result := Rand(length)
		hasDigit := false
		hasUpper := false
		hasLower := false

		for _, char := range result {
			if strings.ContainsRune(Numeric, char) {
				hasDigit = true
			}
			if strings.ContainsRune(Uppercase, char) {
				hasUpper = true
			}
			if strings.ContainsRune(Lowercase, char) {
				hasLower = true
			}
		}

		// With 1000 characters, we should see all character types
		if !hasDigit || !hasUpper || !hasLower {
			t.Errorf("Rand(%d) should contain digits, uppercase, and lowercase. Got digit=%v, upper=%v, lower=%v",
				length, hasDigit, hasUpper, hasLower)
		}
	})

	t.Run("RandNumeric only uses digits", func(t *testing.T) {
		result := RandNumeric(length)
		for _, char := range result {
			if char < '0' || char > '9' {
				t.Errorf("RandNumeric contains non-numeric character: %q", char)
			}
		}
	})

	t.Run("RandAlphabetic only uses letters", func(t *testing.T) {
		result := RandAlphabetic(length)
		for _, char := range result {
			if strings.ContainsRune(Numeric, char) {
				t.Errorf("RandAlphabetic contains numeric character: %q", char)
			}
			if !strings.ContainsRune(Alphabetic, char) {
				t.Errorf("RandAlphabetic contains non-alphabetic character: %q", char)
			}
		}
	})
}

// Benchmark tests
func BenchmarkRand(b *testing.B) {
	for b.Loop() {
		Rand(32)
	}
}

func BenchmarkRandNumeric(b *testing.B) {
	for b.Loop() {
		RandNumeric(32)
	}
}

func BenchmarkRandAlphabetic(b *testing.B) {
	for b.Loop() {
		RandAlphabetic(32)
	}
}
