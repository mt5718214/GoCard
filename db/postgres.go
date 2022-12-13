package db

import (
	"database/sql"
	"fmt"
	sqlcDb "gocard/db/sqlc"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var SqlDB *sql.DB
var Queries *sqlcDb.Queries

func init() {
	if err := godotenv.Load(".env", ".env.test"); err != nil {
		log.Print("Error loading all .env file")
	}
	db := os.Getenv("db")
	dbHost := os.Getenv("dbHost")
	dbUser := os.Getenv("dbUser")
	dbPassword := os.Getenv("dbPassword")
	dbSource := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", dbUser, dbPassword, dbHost, db)

	// db connection
	var err error
	SqlDB, err = sql.Open("postgres", dbSource)
	if err != nil {
		fmt.Println("DB連線資訊有誤請再次確認")
	}
	if err := SqlDB.Ping(); err != nil {
		fmt.Println("開啟 POSTGRES 連線發生錯誤，原因為：", err.Error())
		log.Fatal("db connect error: ", err.Error())
	}
	Queries = sqlcDb.New(SqlDB)
}
