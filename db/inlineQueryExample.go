package db

import (
	"context"
	"log"

	"github.com/haron1996/numbers/viper"
	"github.com/jackc/pgx"
)

func DeleteNumbersWithWrongPrefixes() {
	config, err := viper.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	dbConfig, err := pgx.ParseConnectionString(config.DBString)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.Connect(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	query := `DELETE FROM mobile_number 
	WHERE NOT (
		mobile_number LIKE '070%' OR
		mobile_number LIKE '071%' OR
		mobile_number LIKE '072%' OR
		mobile_number LIKE '074%' OR
		mobile_number LIKE '079%' OR
		mobile_number LIKE '0757%' OR
		mobile_number LIKE '0758%' OR
		mobile_number LIKE '0759%' OR
		mobile_number LIKE '0768%' OR
		mobile_number LIKE '0769%' OR
		mobile_number LIKE '0110%' OR
		mobile_number LIKE '0111%' OR
		mobile_number LIKE '0112%' OR
		mobile_number LIKE '0113%' OR
		mobile_number LIKE '0114%' OR
		mobile_number LIKE '0115%'
	);`

	_, err = conn.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Records deleted successfully")
}
