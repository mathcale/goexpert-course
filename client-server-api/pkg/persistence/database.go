package persistence

import "database/sql"

type Database struct {
	Connection *sql.DB
}

func NewDatabase(connection *sql.DB) *Database {
	return &Database{
		Connection: connection,
	}
}
