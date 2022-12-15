package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	dbDrive = "postgres"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	var err error

	viper.SetConfigFile("../../.env")
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Fatal("fatal error config file: ", err.Error())
	}

	conn, err := sql.Open(dbDrive, viper.GetString("TEST_DB_SOURCE"))
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
