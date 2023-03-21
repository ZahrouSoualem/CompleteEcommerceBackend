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
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/ecommerce?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("We Cannot Connect to the database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
