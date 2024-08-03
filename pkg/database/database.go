package database

import (
	"database/sql"
)

func New(driver string, source string) (database *sql.DB) {
	database, err := sql.Open(driver, source)
	if err != nil {
		panic(err)
	}

	return
}
