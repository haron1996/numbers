package random

import (
	"os"
	"slices"
	"testing"

	"github.com/haron1996/numbers/remove"
)

func TestSelectRandom(t *testing.T) {

	tempCSV := "test.csv"

	defer remove.RemoveTempCSV(tempCSV, t)

	content := "0700123456\n0711123456\n0722123456\n0797476043\n0796701137\n0704642532\n0715332968\n0799386242\n0712918202\n0719296473\n0717608911\n0727888685\n0715150741\n0703659288\n0716544250\n0711401235\n0719947394\n0729461293\n0722860315\n0742012673\n0716437675\n0799273468\n0716379562"

	err := os.WriteFile(tempCSV, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Error creating temporary CSV file: %v", err)
	}

	numRandomRecords := 10

	result, err := SelectRandom(tempCSV, numRandomRecords)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result) != numRandomRecords {
		t.Errorf("Expected %d but got %d", numRandomRecords, len(result))
	}

	expectedRecords := []string{
		"0700123456",
		"0711123456",
		"0722123456",
		"0797476043",
		"0796701137",
		"0704642532",
		"0715332968",
		"0799386242",
		"0712918202",
		"0719296473",
		"0717608911",
		"0727888685",
		"0715150741",
		"0703659288",
		"0716544250",
		"0711401235",
		"0719947394",
		"0729461293",
		"0722860315",
		"0742012673",
		"0716437675",
		"0799273468",
		"0716379562",
	}

	for _, r := range result {
		if !slices.Contains(expectedRecords, r) {
			t.Errorf("%s does not exist in expected records", r)
		}
	}
}
