package main

import (
	"database/sql"
	"log"

	"github.com/zahrou/ecommerce/api"
	db "github.com/zahrou/ecommerce/db/sqlc"

	// by this package now our complier will know our specific database engine
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/ecommerce?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("We Cannot Connect to the database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("We cannot connect to the database", err)
	}
}
