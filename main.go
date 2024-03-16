package main

import (
	"log"

	"github.com/haron1996/numbers/generate"
)

func main() {

	c := 500000
	csvFile := "numbers.csv"
	// numRandomRecords := 10

	if err := generate.GenAndSave(csvFile, c); err != nil {
		log.Println(err)
		return
	}

	// if err := clean.Clean(csvFile); err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// total, err := count.CountTotalNumbers(csvFile)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// log.Printf("Total numbers: %d", total)

	// records, err := random.SelectRandom(csvFile, numRandomRecords)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// log.Printf("%d records selected: %v", len(records), records)

	//find.FindPhoneNumber(csvFile, "0718448461")
	//search.SearchPhoneNumber(csvFile, "0724339139")
}
