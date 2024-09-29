package main

import (
	"database/sql"
	"github.com/DreamCreatives/simplebank/api"
	db "github.com/DreamCreatives/simplebank/db/sqlc"
	"github.com/DreamCreatives/simplebank/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatalf("Cannot start server. Error: %v", err.Error())
	}
}
