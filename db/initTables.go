package db

import (
	"database/sql"
)

func initTables(db *sql.DB) {
	userTable(db)
}

func userTable(db *sql.DB) {
	query := `CREATE TABLE usuarios (
		id SERIAL PRIMARY KEY,
		nombre VARCHAR(50)
	);`

	_, _ = db.Exec(query)
}
