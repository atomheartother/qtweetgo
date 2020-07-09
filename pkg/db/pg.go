package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Pg is the single driver for this postgres connection
var Pg *pgx.Conn

// Init initializes the database connection
func Init() error {
	conn, err := pgx.Connect(context.Background(), "")
	if err != nil {
		return err
	}
	Pg = conn
	return nil
}

// Close closes the database connection
func Close() error {
	return Pg.Close(context.Background())
}
