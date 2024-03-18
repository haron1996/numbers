package generate

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func GenAndSave(csvFile string, count int) error {
	file, err := os.OpenFile(csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	for i := 0; i < count; i++ {

		phoneNumber := generateAlternatePhoneNumber()

		if len(phoneNumber) != 10 {
			continue
		}

		err := writer.Write([]string{phoneNumber})
		if err != nil {
			return fmt.Errorf("failed to write to CSV: %v", err)
		}

		log.Printf("%s Saved to CSV", phoneNumber)
	}

	log.Println("Numbers Inserted to CSV Successfully")

	return nil
}

func generatePhoneNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prefixes := []string{"070", "071", "072", "074", "079", "076"}
	randomPrefix := prefixes[r.Intn(len(prefixes))]
	randomSuffix := 100000000 + r.Intn(900000000)
	randomSuffixStr := fmt.Sprintf("%09d", randomSuffix)
	phoneNumber := fmt.Sprintf("%s%s", randomPrefix, randomSuffixStr)
	phoneNumber = phoneNumber[:10]
	return phoneNumber
}
