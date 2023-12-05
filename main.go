package main

import (
	"l2/mux/controllers"
	"l2/mux/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	repository.InitialTestData()

	r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
