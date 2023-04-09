package main

import (
	"database/sql"
	"log"

	"github.com/zahrou/ecommerce/api"
	db "github.com/zahrou/ecommerce/db/sqlc"
	"github.com/zahrou/ecommerce/util"

	// by this package now our complier will know our specific database engine
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load the configuration")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	//defer conn.Close()

	if err != nil {
		log.Fatal("We Cannot Connect to the database")
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("We Cannot Connect to the database")
	}

	// this to start the server on specific address and port number
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("We cannot connect to the database", err)
	}
}
