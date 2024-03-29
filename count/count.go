package count

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CountTotalNumbers(csvName string) (int, error) {

	file, err := os.Open(csvName)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %v", err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()

	log.Printf("Total Rows: %d", len(fileLines))

	return len(fileLines), nil
}
