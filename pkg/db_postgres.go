package pkg

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetPgConn() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "host=localhost port=7000 user=postgres password=root dbname=db_pos_app sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, err
}
