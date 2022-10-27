package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"restaurant-browser/env"
	"time"
)

type PostgresSQL struct {
	ctx        context.Context
	config     env.EnvApp
	connection *sql.DB
}

// Max seconds for retry a database connection
const DB_CONNECTION_TIMEOUT = 10

func NewPostgresSQL(ctx context.Context, ec env.EnvApp) Database {
	return &PostgresSQL{ctx: ctx, config: ec, connection: nil}
}

func (p *PostgresSQL) ping(db *sql.DB) error {
	err := db.PingContext(p.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresSQL) ConnectDB() error {
	counts := 0

	for {
		db, err := p.connect()

		if try(err, db, &counts) == nil {
			p.connection = db
			return nil
		}
		continue
	}
}

func (p *PostgresSQL) CloseDB() {
	if p.connection != nil {
		_ = p.connection.Close()
	}
}

func (p *PostgresSQL) GetConnection() *sql.DB {
	return p.connection
}

func (p *PostgresSQL) connect() (*sql.DB, error) {
	dataSN := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		p.config.DB_USER,
		p.config.DB_PASSWORD,
		p.config.DB_HOST,
		p.config.DB_PORT,
		p.config.DB_NAME,
	)

	db, err := sql.Open("postgres", dataSN)
	if err != nil {
		return nil, err
	}

	if err = p.ping(db); err != nil {
		return nil, err
	}

	return db, nil
}

// Try db connection
func try(err error, db *sql.DB, counts *int) error {
	if err != nil {
		// increase counter
		fmt.Fprintf(os.Stdout, "Trying to connect with database: %s\n", err.Error())
		*counts++

		// can't connect with the database
		if *counts > DB_CONNECTION_TIMEOUT {
			fmt.Fprintf(os.Stdout, "Can't connect with the database: %s\n", err.Error())
		}

		// log and try again
		fmt.Fprintf(os.Stdout, "Backing off for a second: %s\n", err.Error())
		time.Sleep(time.Second)
	}

	return err
}
