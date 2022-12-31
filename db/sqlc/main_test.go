package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	// sql driver necessary to connct to postgres
	"github.com/Diego-Mike/hardcore-backend-golang/util"
	_ "github.com/lib/pq"
)

// IMPORTANT --> we are using unit test
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("Can't load env variables", err)
	}

	// open connection to db usign built in package database/sql
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Can't connecto to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
