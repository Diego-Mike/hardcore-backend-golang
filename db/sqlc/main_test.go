package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	// sql driver necessary to connct to postgres
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:diegoelmono@localhost:5432/simple_bank?sslmode=disable"
)

// IMPORTANT --> we are using unit test
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	var err error

	// open connection to db usign built in package database/sql
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Can't connecto to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
