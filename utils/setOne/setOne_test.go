package setOne

import (
	"regexp"
	"testing"
)

func TestSetOne(t *testing.T) {

	for i := 0; i < 1000000; i++ {
		phoneNumber := GenerateStartingWith0700710720740And79()

		if len(phoneNumber) != 10 {
			t.Errorf("Generated phone number %s is not 10 digits long", phoneNumber)
		}

		validPrefixes := []string{"070", "071", "072", "074", "079"}

		isValidPrefix := false

		for _, prefix := range validPrefixes {
			if phoneNumber[:3] == prefix {
				isValidPrefix = true
				break
			}
		}

		if !isValidPrefix {
			t.Errorf("Generated phone number %s does not start with a valid prefix", phoneNumber)
		}

		digitRegex := regexp.MustCompile(`^\d+$`)

		if !digitRegex.MatchString(phoneNumber[3:]) {
			t.Errorf("Generated phone number %s contains non-digit characters", phoneNumber)
		}
	}
}
