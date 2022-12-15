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
	dbSource = "postgres://gocard:secret@localhost:5432/gocard_test?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDrive, dbSource)
	if err != nil {
		log.Fatal("DB連線資訊有誤請再次確認: ", err.Error())
	}
	if err := conn.Ping(); err != nil {
		log.Println("開啟 POSTGRES 連線發生錯誤，原因為：", err.Error())
		log.Fatal("db connect error: ", err.Error())
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
