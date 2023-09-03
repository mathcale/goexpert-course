package persistence

import "database/sql"

type Database struct {
	Instance *sql.DB
}

func NewDatabase(instance *sql.DB) Database {
	return Database{
		Instance: instance,
	}
}
