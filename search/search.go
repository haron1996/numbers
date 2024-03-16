package search

import (
	"encoding/csv"
	"log"
	"os"
)

func SearchPhoneNumber(csvFile, number string) (bool, error) {

	file, err := os.Open(csvFile)
	if err != nil {
		log.Println("Error opening file:", err)
		return false, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		log.Println("Error reading CSV:", err)
		return false, err
	}

	searchValue := number

	for _, row := range rows {
		for _, value := range row {
			if value == searchValue {
				log.Printf("%v Found!", row)
				return true, nil
			}
		}
	}

	log.Printf("[%s] Not Found!", searchValue)

	return false, nil
}
