package db

import (
	"database/sql"
	"log"
)

func OpenDb(force bool) *sql.DB {
	connStr := "user=postgres password=admin dbname=golang_tutorial sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	resetDb(db, force)
	initTables(db)

	return db
}

func resetDb(db *sql.DB, force bool) {
	if force {
		query := `DO $$ DECLARE
			r RECORD;
		BEGIN
			FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
					EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' RESTART IDENTITY CASCADE;';
			END LOOP;
		END $$;`

		_, _ = db.Exec(query)
	}
}
