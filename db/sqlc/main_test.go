package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://bank:welcome@localhost:5432/bank?sslmode=disable"
)

var queries *Queries

func TestMain(m *testing.M) {
	pool, err := pgxpool.New(context.Background(), "user=bank password=welcome host=127.0.0.1 port=5432 dbname=bank sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	queries = New(pool)

	os.Exit(m.Run())
}
