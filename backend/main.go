package main

import (
	"log"
	"log/slog"
	"msgr/controller"
	"msgr/database"
	jwthandling "msgr/jwt-handling"
	"net/http"
	"os"

	"msgr/routes"

	"github.com/joho/godotenv"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

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

	var frontendUrl string = os.Getenv("FRONTEND_URL")
	if frontendUrl == "" {
		log.Fatal("variable FRONTEND_URL was not found in .env file")
	}

	var jwtSecret string = os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("variable JWT_SECRET was not found in .env file")
	}

	ctx, conn, err := database.GetConnection(dbUrl)
	if err != nil {
		log.Fatal("failed to connect to db, check if url provided is valid")
	}
	defer conn.Close()

	queries := database.New(conn)

	controller.Initialize(frontendUrl, ctx, queries)
	jwthandling.TokenSecret = []byte(jwtSecret)

	router := routes.CreateRouter()

	slog.Info("started server at port " + port)
	http.ListenAndServe(":"+port, router)
}
