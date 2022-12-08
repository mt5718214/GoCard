package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDrive  = "postgres"
	dbSource = "postgres://gocard:secret@localhost:5432/gocard?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDrive, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err.Error())
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
