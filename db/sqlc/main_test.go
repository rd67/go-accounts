package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rd67/go-accounts/util"
)

var testQueries *Queries

func TestMain(m *testing.M) {
    config, err := util.LoadConfig("../..")
    if err != nil {
        log.Fatal("Error loading config: ", err)
		return
    }

	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
		return
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
