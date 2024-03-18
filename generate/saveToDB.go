package generate

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	sqlc "github.com/haron1996/numbers/db/sqlc"
	"github.com/haron1996/numbers/viper"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GenAndSaveToDB(count int) error {
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

	for i := 0; i < count; i++ {

		phoneNumber := generateAlternatePhoneNumber()

		if len(phoneNumber) != 10 {
			continue
		}

		err = q.InsertMobileNumber(ctx, phoneNumber)
		if err != nil {
			log.Fatal(err)
		}

	}

	log.Println("Insert Successful")

	return nil
}

func generateAlternatePhoneNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Intn(2) == 0 {
		return generate011PhoneNumber()
	}
	return generatePhoneNumber()
}

func generate011PhoneNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomSuffix := 1000000 + r.Intn(9000000)
	phoneNumber := fmt.Sprintf("011%d", randomSuffix)
	if len(phoneNumber) > 10 {
		phoneNumber = phoneNumber[:10]
	}
	return phoneNumber
}
