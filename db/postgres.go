package db

import (
	"database/sql"
	"fmt"
	sqlcDb "gocard/db/sqlc"
	"log"

	_ "github.com/lib/pq"
)

var SqlDB *sql.DB
var Queries *sqlcDb.Queries

func init() {
	// db connection
	var err error
	SqlDB, err = sql.Open("postgres", "user=gocard password=secret dbname=gocard sslmode=disable")
	if err != nil {
		fmt.Println("DB連線資訊有誤請再次確認")
	}
	if err := SqlDB.Ping(); err != nil {
		fmt.Println("開啟 POSTGRES 連線發生錯誤，原因為：", err.Error())
		log.Fatal("db connect error: ", err.Error())
	}
	Queries = sqlcDb.New(SqlDB)
}
