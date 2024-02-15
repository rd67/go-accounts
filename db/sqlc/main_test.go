package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB_DRIVER            = ""
	DB_CONNECTION_STRING = ""
)

const projectDirName = "go-accounts" // change to relevant project name

func loadEnv() {
	// Load variables from .env file
    projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
    currentWorkDirectory, _ := os.Getwd()
    rootPath := projectName.Find([]byte(currentWorkDirectory))

    err := godotenv.Load(string(rootPath) + `/.env`)

    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}

//TODO: *** Duplicate Code remove, improve 
func init() {
	loadEnv()

	DB_DRIVER = os.Getenv("DB_DRIVER")

	//	"${DB_DRIVER}://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):${DB_PORT})/$(DB_NAME)"
	DB_CONNECTION_STRING = fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

}

var testQueries *Queries

func TestMain(m *testing.M) {
	fmt.Println("env variables loaded DB_DRIVER: ",DB_DRIVER)

	conn, err := sql.Open(DB_DRIVER, DB_CONNECTION_STRING)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}