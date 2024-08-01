package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func New(driver string, source string) (database *sql.DB) {
	database, err := sql.Open(driver, source)
	if err != nil {
		panic(err)
	}

	return
}
