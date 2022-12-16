package db

import (
	"database/sql"
	"fmt"
	sqlcDb "gocard/db/sqlc"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var SqlDB *sql.DB
var Queries *sqlcDb.Queries

func init() {
	var err error

	viper.SetConfigFile("../.env")
	// func ReadInConfig Find and read the config file
	if err = viper.ReadInConfig(); err != nil {
		log.Fatal("fatal error config file: ", err.Error())
	}

	// db connection
	SqlDB, err = sql.Open("postgres", viper.GetString("DB_SOURCE"))
	if err != nil {
		fmt.Println("DB連線資訊有誤請再次確認")
	}
	if err := SqlDB.Ping(); err != nil {
		fmt.Println("開啟 POSTGRES 連線發生錯誤，原因為：", err.Error())
		log.Fatal("db connect error: ", err.Error())
	}
	Queries = sqlcDb.New(SqlDB)
}
