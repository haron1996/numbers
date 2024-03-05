package remove

import (
	"os"
	"testing"
)

func RemoveTempCSV(csvFile string, t *testing.T) {
	err := os.Remove(csvFile)
	if err != nil {
		t.Fatalf("Error removing csv file")
	}
}
