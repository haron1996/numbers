package generate

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/haron1996/numbers/remove"
)

func TestGenAndSave(t *testing.T) {
	tempCSV := "test.csv"

	defer remove.RemoveTempCSV(tempCSV, t)

	count := 5

	err := GenAndSave(tempCSV, count)
	if err != nil {
		t.Fatalf("Error generating and saving phone numbers: %v", err)
	}

	file, err := os.Open(tempCSV)
	if err != nil {
		t.Fatalf("Error opening CSV file: %v", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	readContent, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Error reading CSV file: %v", err)
	}

	if len(readContent) != count {
		t.Errorf("Number of generated phone numbers does not match expected count")
	}

	for _, record := range readContent {
		if len(record) != 1 || len(record[0]) != 10 {
			t.Errorf("Invalid phone number format: %v", record)
		}
	}
}
