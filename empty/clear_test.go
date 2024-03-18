package empty

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/haron1996/numbers/remove"
)

func TestEmptyCSV(t *testing.T) {
	tempCSV := "test.csv"

	defer remove.RemoveTempCSV(tempCSV, t)

	content := "0700123456\n0711123456\n0722123456"

	err := os.WriteFile(tempCSV, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Error creating temporary CSV file: %v", err)
	}

	EmptyCSV(tempCSV)

	file, err := os.Open("test.csv")
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Erro reading csv: %v", err)
	}

	if len(records) > 0 {
		t.Errorf("CSV not empty after clearing")
	}
}
