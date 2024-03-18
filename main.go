package main

import (
	"github.com/haron1996/numbers/clean"
	"github.com/haron1996/numbers/count"
	"github.com/haron1996/numbers/db"
	"github.com/haron1996/numbers/empty"
	"github.com/haron1996/numbers/generate"
)

func main() {
	for i := 0; i < 30; i++ {
		generate.GenAndSave("random.csv", 5000000)
		clean.Clean("random.csv")
		count.CountTotalNumbers("random.csv")
		db.InsertMobileNumbersToDB_Tx("random.csv")
		db.DeleteNumbersWithWrongPrefixes()
		db.GetAllMobileNumbers()
		empty.EmptyCSV("random.csv")
	}
}
