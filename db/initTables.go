package db

import (
	"database/sql"
)

func initTables(db *sql.DB) {
	userTable(db)
	taskTable(db)
}

func userTable(db *sql.DB) {
	query := `CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		firstName VARCHAR(50) NOT NULL,
		lastName VARCHAR(50) NOT NULL
	);`

	_, _ = db.Exec(query)
}

func taskTable(db *sql.DB) {
	query := `CREATE TABLE tasks (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50),
		userId INTEGER REFERENCES users(id)
	);`

	_, _ = db.Exec(query)
}
