package database

type Database interface {
	ConnectDB() error
	CloseDB()
}
