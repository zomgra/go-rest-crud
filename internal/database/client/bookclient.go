package client

import (
	"golangCrud/internal/database"
	"golangCrud/pkg/entities"
	"log"
)

func AddBook(b entities.Book) int {
	log.Printf("%+v", b)
	query := `INSERT INTO books (title, isnbn) VALUES($1,$2) RETURNING id`
	var id int
	err := database.Instance.QueryRow(query, b.Title, b.Isbn).Scan(&id)
	if err != nil {
		log.Fatal("problem with insert book: ", err)
	}
	return id
}
func GetAllBooks() []entities.Book {
	query := `SELECT * FROM "books"`
	rows, err := database.Instance.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var books []entities.Book
	for rows.Next() {
		var book entities.Book
		rows.Scan(&book.Id, &book.Title, &book.AuthorId, &book.Isbn)
		if book.Id == "" {
			break
		}
		books = append(books, book)
	}
	return books
}
func UpdateBook(id string, b entities.Book) {
	query := `UPDATE books
	SET title=$2, isnbn=$3
	WHERE id=$1`
	log.Print(b)
	_, err := database.Instance.Exec(query, id, b.Title, b.Isbn)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteBook(id string) {
	query := `DELETE FROM books WHERE id=$1`
	_, err := database.Instance.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
}

func GetBookById(id string) entities.Book {
	query := `SELECT id, title, isnbn FROM books WHERE id=$1 LIMIT 1`
	var book entities.Book
	err := database.Instance.QueryRow(query, id).Scan(&book.Id, &book.Title, &book.Isbn)
	if err != nil {
		log.Fatal(err)
	}
	return book
}
