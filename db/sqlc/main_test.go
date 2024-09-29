package db

import (
	"database/sql"
	"github.com/DreamCreatives/simplebank/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot get config for tests", err)
	}

	testDB, err = sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
