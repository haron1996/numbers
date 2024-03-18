package setTwo

import (
	"testing"
)

func TestGenerateRandomNumber(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		result := GenerateStartingWith0757075807590768And0769()

		if len(result) != 10 {
			t.Errorf("Generated number %s has incorrect length", result)
		}

		validPrefix := false

		validPrefixes := []string{"0757", "0758", "0759", "0768", "0769"}
		for _, prefix := range validPrefixes {
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
