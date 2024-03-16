package generate

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func GenAndSave(csvFile string, count int) error {

	for i := 0; i < count; i++ {

		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		prefixes := []string{"070", "071", "072", "074", "079", "076"}

		randomPrefix := prefixes[r.Intn(len(prefixes))]

		randomSuffix := 100000000 + r.Intn(900000000)

		phoneNumber := fmt.Sprintf("%s%d", randomPrefix, randomSuffix)[0:10]

		if len(phoneNumber) < 10 {
			continue
		}

		file, err := os.OpenFile(csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open CSV file: %v", err)
		}

		defer file.Close()

		writer := csv.NewWriter(file)

		defer writer.Flush()

		err = writer.Write([]string{phoneNumber})
		if err != nil {
			return fmt.Errorf("failed to write to CSV: %v", err)
		}
	}

	return nil
}
