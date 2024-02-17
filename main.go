package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rd67/go-accounts/api"
	db "github.com/rd67/go-accounts/db/sqlc"
)

var (
	DB_DRIVER            = ""
	DB_CONNECTION_STRING = ""
)

func init() {
	// Load variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	DB_DRIVER = os.Getenv("DB_DRIVER")

	//	"${DB_DRIVER}://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):${DB_PORT})/$(DB_NAME)"
	DB_CONNECTION_STRING = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

}

func main() {
	// Print variables
	fmt.Println("env variables loaded....")

	conn, err := sql.Open(DB_DRIVER, DB_CONNECTION_STRING)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(":"+os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal("Error running server", err)
	}
}
