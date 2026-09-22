package database

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	Connection       *sql.DB
	connectionString string
	connected        bool
}

func NewDatabseConnection(connectionString string) DatabaseConnection {
	connection := DatabaseConnection{connectionString: connectionString, connected: false}
	return connection
}

func (r *DatabaseConnection) Connect() error {
	var err error
	if r.connected {
		return errors.New("aleady connected")
	}
	r.connected = true
	r.Connection, err = sql.Open("postgres", r.connectionString)
	return err
}
