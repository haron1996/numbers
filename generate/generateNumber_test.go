package generate

import "testing"

func TestGenerate(t *testing.T) {

	for i := 0; i < 1000000; i++ {
		phoneNumber := generatePhoneNumber()
		if len(phoneNumber) != 10 {
			t.Errorf("Generated phone number %s does not have 10 characters", phoneNumber)
		}
	}
}
