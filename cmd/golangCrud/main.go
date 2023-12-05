package main

import (
	"golangCrud/internal/database"
	"golangCrud/internal/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	r := routes.NewRoute()

	opts := database.DbOptions{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_DATABASE"),
	}
	database.CreateConnection(&opts)
	log.Fatal(http.ListenAndServe(":8000", r))
}
