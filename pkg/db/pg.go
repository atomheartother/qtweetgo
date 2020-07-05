package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// Pg is the single driver for this postgres connection
var Pg *pgx.Conn

func init() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Pg = conn
}
