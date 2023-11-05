package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type DatabaseInterface interface {
	Initialize()
	GetPool() *pgxpool.Pool
}

type Database struct {
	credentials string
}

var _ DatabaseInterface = &Database{}

func (d *Database) GetPool() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), d.credentials)

	if err != nil {
		log.Fatal(err)
	}

	return pool
}

func (d *Database) Initialize() {
	d.credentials = "user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable"
}
