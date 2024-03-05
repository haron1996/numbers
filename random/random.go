package random

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func SelectRandom(csvFile string, numRandomRecords int) ([]string, error) {

	file, err := os.Open(csvFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	numRecords := len(records)

	if numRecords == 0 {
		return nil, errors.New("CSV file is empty")
	}

	randomIndices := generateRandomIndices(numRecords, numRandomRecords)

	var selectedRecords []string

	for _, index := range randomIndices {
		if index >= numRecords {
			return nil, fmt.Errorf("generated index out of bounds: %d (total records: %d)", index, numRecords)
		}

		record := records[index]
		selectedRecords = append(selectedRecords, record...)
	}

	return selectedRecords, nil
}

func generateRandomIndices(maxIndex, count int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	indices := make([]int, count)

	for i := 0; i < count; i++ {
		// Generate a random index
		index := r.Intn(maxIndex)

		// Check for uniqueness
		for contains(indices, index) {
			index = r.Intn(maxIndex)
		}

		indices[i] = index
	}

	return indices
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
