package setThree

import (
	"regexp"
	"testing"
)

func TestSetThree(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		phoneNumber := GenerateStartingWith01100111011201130114And0115()

		if len(phoneNumber) != 10 {
			t.Errorf("Generated phone number %s is not 10 digits long", phoneNumber)
		}

		validPrefixes := []string{"0110", "0111", "0112", "0113", "0114", "0115"}

		isValidPrefix := false

		for _, prefix := range validPrefixes {
			if phoneNumber[:4] == prefix {
				isValidPrefix = true
				break
			}
		}

		if !isValidPrefix {
			t.Errorf("Generated phone number %s does not start with a valid prefix", phoneNumber)
		}

		digitRegex := regexp.MustCompile(`^\d+$`)

		if !digitRegex.MatchString(phoneNumber[4:]) {
			t.Errorf("Generated phone number %s contains non-digit characters", phoneNumber)
		}
	}
}
