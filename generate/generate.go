package generate

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/haron1996/numbers/utils/alternate"
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

		phoneNumber := alternate.GeneratePhoneNumber()

		if len(phoneNumber) != 10 {
			continue
		}

		err := writer.Write([]string{phoneNumber})
		if err != nil {
			return fmt.Errorf("failed to write to CSV: %v", err)
		}

	}

	log.Println("Numbers Inserted to CSV Successfully")

	return nil
}
