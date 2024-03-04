package funcs

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"time"
)

func SelectRandom(csvFile string, numRandomRecords int) error {

	file, err := os.Open(csvFile)
	if err != nil {
		log.Println("could not open csv file")
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Println("could not read csv file")
		return err
	}

	numRecords := len(records)

	if numRecords == 0 {
		log.Println("CSV file is empty.")
		return err
	}

	randomIndices := generateRandomIndices(numRecords, numRandomRecords)

	for _, index := range randomIndices {
		record := records[index]
		log.Println("Random Record:", record)
	}

	return nil
}

func generateRandomIndices(max int, count int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomIndices := make([]int, count)

	perm := r.Perm(max)

	for i := 0; i < count; i++ {
		randomIndices[i] = perm[i]
	}

	return randomIndices
}
