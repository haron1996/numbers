package funcs

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func Clean(csvFile string) error {
	startTime := time.Now()

	file, err := os.Open(csvFile)
	if err != nil {
		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	uniqueRecords := make(map[string]bool)

	var finalRecords [][]string

	for {

		record, err := reader.Read()

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		recordString := fmt.Sprintf("%v", record)

		if _, exists := uniqueRecords[recordString]; !exists {

			uniqueRecords[recordString] = true
			finalRecords = append(finalRecords, record)
		}
	}

	if err := os.Remove(csvFile); err != nil {
		log.Fatalf("failed to remove csv file: %v", err)
	}

	file, err = os.Create(csvFile)
	if err != nil {
		log.Fatalf("failed to create csv file: %v", err)
	}

	writer := csv.NewWriter(file)

	err = writer.WriteAll(finalRecords)
	if err != nil {
		return err
	}

	writer.Flush()

	log.Printf("Executed in %f seconds", time.Since(startTime).Seconds())

	return nil
}
