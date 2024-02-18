package main

import (
	"database/sql"

	"log"

	_ "github.com/lib/pq"
	"github.com/rd67/go-accounts/api"
	db "github.com/rd67/go-accounts/db/sqlc"
	util "github.com/rd67/go-accounts/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config: ", err)
		return
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
		return
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(":" + config.PORT)
	if err != nil {
		log.Fatal("Error running server", err)
	}
}
