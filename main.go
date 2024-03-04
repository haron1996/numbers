package main

import (
	"log"

	"github.com/haron1996/go-sms/funcs"
)

func main() {

	count := 500000
	csvFile := "numbers.csv"
	// numRandomRecords := 10

	if err := funcs.GenAndSave(csvFile, count); err != nil {
		log.Println(err)
		return
	}

	// funcs.Clean(csvFile)

	total, err := funcs.CountTotalNumbers(csvFile)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Total Numbers: %d", total)

	// if err := funcs.SelectRandom(csvFile, numRandomRecords); err != nil {
	// 	log.Println(err)
	// }
}
