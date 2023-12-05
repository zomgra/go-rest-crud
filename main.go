package main

import (
	"l2/mux/controllers"
	"l2/mux/database"
	"l2/mux/repository"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	r := mux.NewRouter()
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Error loading app.env file")
	}
	repository.InitialTestData()

	r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")

	opts := database.DbOptions{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_DATABASE"),
	}
	database.CreateConnection(&opts)
	log.Fatal(http.ListenAndServe(":8000", r))
}
