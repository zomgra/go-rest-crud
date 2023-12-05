package routes

import (
	"golangCrud/internal/controllers"
	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func NewRoute() *mux.Router {
	r := mux.NewRouter()
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Error loading app.env file")
	}

	r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")

	return r
}
