package main

import (
	"log"

	"github.com/haron1996/numbers/clean"
	"github.com/haron1996/numbers/count"
	"github.com/haron1996/numbers/generate"
	"github.com/haron1996/numbers/random"
)

func main() {

	c := 500000
	csvFile := "numbers.csv"
	numRandomRecords := 10

	if err := generate.GenAndSave(csvFile, c); err != nil {
		log.Println(err)
		return
	}

	if err := clean.Clean(csvFile); err != nil {
		log.Println(err)
		return
	}

	total, err := count.CountTotalNumbers(csvFile)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Total numbers: %d", total)

	records, err := random.SelectRandom(csvFile, numRandomRecords)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%d records selected: %v", len(records), records)
}
