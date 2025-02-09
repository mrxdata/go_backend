package database

import (
	"flutty_messenger/app/queries"
	"github.com/jmoiron/sqlx"
)

const (
	PostgreSQL = 0
)

type Queries struct {
	*queries.UserQueries
}

func OpenDBConnection(dbType int) (*Queries, error) {
	var (
		db  *sqlx.DB
		err error
	)

	if dbType == PostgreSQL {
		db, err = PostgreSQLConnection()
	}

	if err != nil {
		return nil, err
	}

	return &Queries{
		UserQueries: &queries.UserQueries{DB: db},
	}, nil
}
