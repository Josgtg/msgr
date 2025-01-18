package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file was not found!")
	}

	var port string = os.Getenv("PORT")
	if port == "" {
		log.Fatal("variable PORT was not found in .env file")
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, router *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
