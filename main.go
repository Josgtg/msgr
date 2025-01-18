package main

import (
	"log"
	"msgr/database"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file was not found!")
	}

	var port string = os.Getenv("PORT")
	if port == "" {
		log.Fatal("variable PORT was not found in .env file")
	}

	var dbUrl string = os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("variable DB_URL was not found in .env file")
	}

	ctx, conn, err := database.GetConnection(dbUrl)
	if err != nil {
		log.Fatal("failed to connect to db, check if url provided is valid")
	}
	defer conn.Close(ctx)

	_ = database.New(conn)
}
