package count

import (
	"os"
	"testing"

	"github.com/haron1996/numbers/remove"
)

func TestCountTotalNumbers(t *testing.T) {

	tempCSV := "test.csv"

	defer remove.RemoveTempCSV(tempCSV, t)

	content := "0700123456\n0711123456\n0722123456"

	err := os.WriteFile(tempCSV, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Error creating temporary CSV file: %v", err)
	}

	result, err := CountTotalNumbers(tempCSV)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := 3

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
