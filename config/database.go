package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	conStr := "host=localhost port=9000 user=postgres password=123456 dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", conStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	log.Println("Success for database connected")
}
