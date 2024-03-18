package alternate

import (
	"testing"
)

func TestGeneratePhoneNumber(t *testing.T) {
	specialCount := 0
	regularCount := 0
	anotherCount := 0

	for i := 0; i < 1000000; i++ {
		phoneNumber := GeneratePhoneNumber()
		if len(phoneNumber) != 10 {
			t.Errorf("Generated phone number %s does not have 10 characters", phoneNumber)
		}
		switch phoneNumber[:3] {
		case "011":
			specialCount++
		case "070", "071", "072", "074", "079", "076":
			regularCount++
		default:
			anotherCount++
		}
	}

	if specialCount == 0 || regularCount == 0 || anotherCount == 0 {
		t.Errorf("Not alternating properly. Special Count: %d, Regular Count: %d, Another Count: %d", specialCount, regularCount, anotherCount)
	}
}
