package db

import (
	"database/sql"
	"log"
)

func OpenDb() {
	connStr := "user=postgres password=admin dbname=golang_tutorial sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	initTables(db)

	defer db.Close()

}
