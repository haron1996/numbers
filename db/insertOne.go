package db

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	sqlc "github.com/haron1996/numbers/db/sqlc"
	"github.com/haron1996/numbers/viper"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertMobileNumbersToDB(csvFile string) {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	config, err := viper.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, config.DBString)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	pool.Config().MinConns = 100000000

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	q := sqlc.New(pool)

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}

		if err := q.InsertMobileNumber(ctx, record[0]); err != nil {
			log.Fatal(err)
		}
	}
}
