package pgdb

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(url string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
