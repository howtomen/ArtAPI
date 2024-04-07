package db

import (
	"context"
	"fmt"
	// "log"
	"os"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	Client *sqlx.DB
}

func NewDbConnection() (*Database, error) {
	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)
	db, err := sqlx.Connect("postgres", connectionStr)
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to the database: %w", err)
	}

	return &Database{
		Client: db,
	}, nil 
}

//Adding ping function so that later we can health check the conn to DB easily
func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}