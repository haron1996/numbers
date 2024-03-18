package generate

import "testing"

func TestGenerateAlternatePhoneNumber(t *testing.T) {
	specialCount := 0
	regularCount := 0

	for i := 0; i < 1000000; i++ {
		phoneNumber := generateAlternatePhoneNumber()
		if len(phoneNumber) != 10 {
			t.Errorf("Generated phone number %s does not have 10 characters", phoneNumber)
		}
		if phoneNumber[:3] == "011" {
			specialCount++
		} else {
			regularCount++
		}
	}

	if specialCount == 0 || regularCount == 0 {
		t.Errorf("Not alternating properly. Special Count: %d, Regular Count: %d", specialCount, regularCount)
	}
}
