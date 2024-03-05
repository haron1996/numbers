package clean

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"

	"github.com/haron1996/numbers/remove"
)

func createTestCSV(filePath string, content [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	err = writer.WriteAll(content)
	if err != nil {
		return err
	}

	return nil
}

func TestClean(t *testing.T) {

	tempCSV := "test.csv"

	defer remove.RemoveTempCSV(tempCSV, t)

	initialContent := [][]string{
		{"1", "0700123456"},
		{"2", "0711123456"},
		{"3", "0722123456"},
		{"1", "0700123456"},
	}

	expectedContent := [][]string{
		{"1", "0700123456"},
		{"2", "0711123456"},
		{"3", "0722123456"},
	}

	err := createTestCSV(tempCSV, initialContent)
	if err != nil {
		t.Fatalf("Error creating temporary CSV file: %v", err)
	}

	err = Clean(tempCSV)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	file, err := os.Open(tempCSV)
	if err != nil {
		t.Fatalf("Error opening cleaned CSV file: %v", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	readContent, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Error reading cleaned CSV file: %v", err)
	}

	if !reflect.DeepEqual(expectedContent, readContent) {
		t.Errorf("Cleaned content does not match expected value")
	}
}
