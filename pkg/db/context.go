package db

import "github.com/jmoiron/sqlx"

type Context interface {
}

type SQLXDB struct {
	dbConnection *sqlx.DB
}

func NewSQLXDB(addr string, opts ...Option) (Context, error) {
	options := options{
		cache: true,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	conn, err := open(addr)
	if err != nil {
		return nil, err
	}

	db := &SQLXDB{
		dbConnection: conn,
	}

	return db, nil
}

type options struct {
	cache bool
}

type Option interface {
	apply(*options)
}
