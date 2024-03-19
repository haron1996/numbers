package zeroSevenTen

import (
	"testing"
)

func TestSevenTen(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		result := Generate0710()

		if len(result) != 10 {
			t.Errorf("Generated number %s has incorrect length", result)
		}

		validPrefix := false

		validPrefixe := []string{"0710"}

		for _, prefix := range validPrefixe {
			if result[:4] == prefix {
				validPrefix = true
				break
			}
		}

		if !validPrefix {
			t.Errorf("Generated number %s doesn't start with valid prefix", result)
		}

		for _, digit := range result[4:] {
			if digit < '0' || digit > '9' {
				t.Errorf("Generated number %s contains non-numeric characters", result)
			}
		}
	}
}
