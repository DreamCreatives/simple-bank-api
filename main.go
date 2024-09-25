package main

import (
	"database/sql"
	"fmt"

	"github.com/DreamCreatives/simplebank/api"
	db "github.com/DreamCreatives/simplebank/db/sqlc"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:postgres@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0"
	serverPort    = 8080
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	addr := fmt.Sprintf("%v:%v", serverAddress, serverPort)
	err = server.Start(addr)

	if err != nil {
		log.Fatalf("Cannot start server. Error: %v", err.Error())
	}
}
