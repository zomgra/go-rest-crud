package controllers

import (
	"encoding/json"
	"errors"
	"l2/mux/database"
	"l2/mux/entities"
	"l2/mux/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for _, item := range repository.Books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	json.NewEncoder(w).Encode(&entities.Book{})
}
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.GetAllBooks())
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for i, book := range repository.Books {
		if book.Id == params["id"] {
			newBook := entities.Book{Id: book.Id, Author: book.Author}
			err := json.NewDecoder(r.Body).Decode(&newBook)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			repository.Books[i] = newBook
			json.NewEncoder(w).Encode(newBook)
			return
		}
	}
	json.NewEncoder(w).Encode(errors.New("book not found"))
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
	for i, book := range repository.Books {
		if book.Id == params["id"] {
			// Remove the book with the matching ID by slicing the slice
			repository.Books = append(repository.Books[:i], repository.Books[i+1:]...)
			return
		}
	}
}
