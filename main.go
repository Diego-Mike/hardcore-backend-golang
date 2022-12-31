package main

import (
	"database/sql"
	"log"

	"github.com/Diego-Mike/hardcore-backend-golang/api"
	db "github.com/Diego-Mike/hardcore-backend-golang/db/sqlc"
	"github.com/Diego-Mike/hardcore-backend-golang/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Can't load env variables", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Can't connecto to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("can't create server", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Can't start server:", err)
	}

}
