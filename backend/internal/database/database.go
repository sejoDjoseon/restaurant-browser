package database

import "database/sql"

type Database interface {
	ConnectDB() error
	CloseDB()
	GetConnection() *sql.DB
}
