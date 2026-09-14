package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	Connection       *sql.DB
	ConnectionString string
}

func (r *DatabaseConnection) Connect() error {
	var err error
	r.Connection, err = sql.Open("postgres", r.ConnectionString)
	return err
}
