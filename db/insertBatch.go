package db

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	sqlc "github.com/haron1996/numbers/db/sqlc"
	"github.com/haron1996/numbers/viper"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertMobileNumbersToDB_Tx(csvFile string) {
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

	pool.Config().MinConns = 1000

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	q := sqlc.New(pool)

	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
			return
		}
		if err := tx.Commit(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	batchSize := 10000000

	batch := make([]string, 0, batchSize)

	for {
		record, err := csvReader.Read()

		if err != nil {
			if err == io.EOF {
				if len(batch) > 0 {
					_, err := q.InsertBatchMobileNumbers(ctx, batch)
					if err != nil {
						log.Println("Error inserting batch:", err)
						return
					}
				}
				break
			}
			log.Fatal(err)
		}

		if len(record) == 0 {
			continue
		}

		if len(record[0]) != 10 {
			log.Printf("Long or Short record: %v", record[0])
			continue
		}

		batch = append(batch, strings.TrimSpace(record[0]))

		if len(batch) >= batchSize {
			_, err := q.InsertBatchMobileNumbers(ctx, batch)
			if err != nil {
				log.Println("Error inserting batch:", err)
				return
			}
			batch = make([]string, 0, batchSize)
		}
	}

	log.Println("Insert Successful")
}
