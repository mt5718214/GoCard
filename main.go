package main

import (
	db "gocard/db"
	docs "gocard/docs"
	"os"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "gocard API doc"
	docs.SwaggerInfo.Description = "This is goCard. You can visit the GitHub repository at https://github.com/mt5718214/GoCard"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/dev/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	defer db.SqlDB.Close()

	ENV := "DEV"
	if os.Getenv("ENV") == "PROD" {
		ENV = os.Getenv("ENV")
	}
	db.NewDB(ENV, "./")

	server := initRouter()
	// By default it serves on :8080 unless a PORT environment variable was defined.
	// router.Run(":3000") for a hard coded port
	server.Run()
}
