package find

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func FindPhoneNumber(csvFile, numberToFind string) (bool, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Printf("Error: %v", err)
		return false, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Printf("Error: %v", err)
		return false, err
	}

	found := false

	for _, record := range records {
		for _, field := range record {
			if field == numberToFind {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if found {
		fmt.Printf("The value %s exists in the CSV file.\n", numberToFind)
	} else {
		fmt.Printf("The value %s does not exist in the CSV file.\n", numberToFind)
	}

	return found, nil
}
