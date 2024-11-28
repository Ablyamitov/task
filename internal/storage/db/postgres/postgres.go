package postgres

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect(connectionUrl string) (*sqlx.DB, error) {

	db, err := sqlx.Connect("pgx", connectionUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
