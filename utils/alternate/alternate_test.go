package alternate

import (
	"testing"
)

func TestGenerateAlternatePhoneNumber(t *testing.T) {
	setOneCount := 0
	setTwoCount := 0
	setThreeCount := 0
	zeroZevenTenCount := 0

	for i := 0; i < 1000000; i++ {
		phoneNumber := GeneratePhoneNumber()
		if len(phoneNumber) != 10 {
			t.Errorf("Generated phone number %s does not have 10 characters", phoneNumber)
		}

		switch phoneNumber[:3] {
		case "070", "071", "072", "074", "079":
			switch phoneNumber[:4] {
			case "0710":
				zeroZevenTenCount++
			default:
				setOneCount++
			}
		case "075", "076":
			setTwoCount++
		default:
			setThreeCount++
		}
	}

	if setOneCount == 0 || setTwoCount == 0 || setThreeCount == 0 || zeroZevenTenCount == 0 {
		t.Errorf("Not alternating properly. Set One Count: %d, Set Two Count: %d, Set Three Count: %d, Zero Zero Ten Count: %d", setOneCount, setTwoCount, setThreeCount, zeroZevenTenCount)
	}
}
