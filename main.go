package main

import (
	db "gocard/db"
)

func main() {
	defer db.SqlDB.Close()
	server := initRouter()
	// By default it serves on :8080 unless a PORT environment variable was defined.
	// router.Run(":3000") for a hard coded port
	server.Run()
}
