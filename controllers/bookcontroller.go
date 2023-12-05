package controllers

import (
	"encoding/json"
	"l2/mux/database"
	"l2/mux/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book := database.GetBookById(params["id"])
	json.NewEncoder(w).Encode(book)
}
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.GetAllBooks())
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	newBook := entities.Book{}
	json.NewDecoder(r.Body).Decode(&newBook)
	database.UpdateBook(params["id"], newBook)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book entities.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	id := database.AddBook(book)
	json.NewEncoder(w).Encode(id)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	database.DeleteBook(params["id"])
}
