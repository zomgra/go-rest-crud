package controllers

import (
	"encoding/json"
	"golangCrud/internal/database/client"
	"golangCrud/pkg/entities"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book := client.GetBookById(params["id"])
	json.NewEncoder(w).Encode(book)
}
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client.GetAllBooks())
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	newBook := entities.Book{}
	json.NewDecoder(r.Body).Decode(&newBook)
	log.Print(newBook)
	client.UpdateBook(params["id"], newBook)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book entities.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	id := client.AddBook(book)
	json.NewEncoder(w).Encode(id)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	client.DeleteBook(params["id"])
}
