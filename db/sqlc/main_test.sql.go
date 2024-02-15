package db

import (
	"database/sql"
	"testing"
)

const (
	dbDriver = "mysql"
	dbSource = ""
)

var testQueries *Queries

func TestMain(m *testing.M) {
	// conn, err := sql.Open(dbDriver, dbSource)

}