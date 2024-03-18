package db

import (
	"context"
	"log"

	sqlc "github.com/haron1996/numbers/db/sqlc"
	"github.com/haron1996/numbers/viper"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllMobileNumbers() {
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

	q := sqlc.New(pool)

	numbers, err := q.GetAllMobileNumbers(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(len(numbers))
}
