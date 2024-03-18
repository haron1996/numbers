package empty

import (
	"log"
	"os"
)

func EmptyCSV(csvFile string) {
	file, err := os.OpenFile(csvFile, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if err := file.Truncate(0); err != nil {
		log.Fatal(err)
	}

	log.Println("csv cleared successfully")
}
