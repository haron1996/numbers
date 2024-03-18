package generate

import "testing"

func TestGenerate011PhoneNumber(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		phoneNumber := generate011PhoneNumber()
		if len(phoneNumber) != 10 || phoneNumber[:3] != "011" {
			t.Errorf("Generated special phone number %s does not meet the criteria", phoneNumber)
		}
	}
}
