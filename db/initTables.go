package db

import (
	"database/sql"
)

func initTables(db *sql.DB) {
	userTable(db)
}

func userTable(db *sql.DB) {
	query := `CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		firstName VARCHAR(50) NOT NULL,
		lastName VARCHAR(50) NOT NULL
	);`

	_, _ = db.Exec(query)
}
