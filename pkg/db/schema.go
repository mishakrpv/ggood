package db

import (
	"github.com/jmoiron/sqlx"
)

var schema = ` 
`

func open(
	addr string,
) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.MustExec(schema)

	return nil, nil
}
