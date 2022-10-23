package database

import (
	"context"
	"database/sql"
	"fmt"
	"restaurant-browser/env"
)

type PostgresSQL struct {
	ctx        context.Context
	config     env.EnvApp
	connection *sql.DB
}

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
		return err
	}

	if err = p.ping(db); err != nil {
		return err
	}

	p.connection = db
	return nil
}

func (p *PostgresSQL) CloseDB() {
	if p.connection != nil {
		_ = p.connection.Close()
	}
}

func (p *PostgresSQL) GetConnection() *sql.DB {
	return p.connection
}
