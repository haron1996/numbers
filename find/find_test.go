package find

import (
	"os"
	"testing"

	"github.com/haron1996/numbers/remove"
)

func TestFindPhoneNumber(t *testing.T) {
	tempCSV := "test.csv"

	defer remove.RemoveTempCSV(tempCSV, t)

	content := "0718448461\n0713194323\n0794558991"

	err := os.WriteFile(tempCSV, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Error creating temporary CSV file: %v", err)
	}

	var tests = []struct {
		name  string
		input string
		want  bool
	}{

		{"0718448461 should be True", "0718448461", true},
		{"0713194323 should be True", "0713194323", true},
		{"0794558991 should be True", "0794558991", true},
		{"0704237652 should be False", "0704237652", false},
		{"0724339139 should be False", "0724339139", false},
		{"0714701847 should be false", "0714701847", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, err := FindPhoneNumber(tempCSV, tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if found != tt.want {
				t.Errorf("for input %s, got %t, want %t", tt.input, found, tt.want)
			}
		})
	}
}
