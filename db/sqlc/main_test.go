package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	// by this package now our complier will know our specific database engine
	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/ecommerce?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("we cannot connect to the database")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
