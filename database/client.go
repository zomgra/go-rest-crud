package database

import (
	"database/sql"
	"fmt"
	"l2/mux/entities"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var Instance *sql.DB

func CreateConnection(o *DbOptions) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", o.User, o.Password, o.Addr, o.Database)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	createMigrations(connString)
	Instance = db // Add migrations
}

func createMigrations(connString string) {
	m, err := migrate.New("file://database/migrations", connString)
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	m.Up()
}

func AddBook(b entities.Book) int {
	log.Printf("%+v", b)
	query := `INSERT INTO books (title, isnbn) VALUES($1,$2) RETURNING id`
	var id int
	err := Instance.QueryRow(query, b.Title, b.Isnb).Scan(&id)
	if err != nil {
		log.Fatal("problem with insert book: ", err)
	}
	return id
}
func GetAllBooks() []entities.Book {
	query := `SELECT * FROM "books"`
	rows, err := Instance.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var books []entities.Book
	for rows.Next() {
		var book entities.Book
		rows.Scan(&book.Id, &book.Title, &book.AuthorId, &book.Isnb)
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
	_, err := Instance.Exec(query, id, b.Title, b.Isnb)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteBook(id string) {
	query := `DELETE FROM books WHERE id=$1`
	_, err := Instance.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
}

func GetBookById(id string) entities.Book {
	query := `SELECT id, title, isnbn FROM books WHERE id=$1 LIMIT 1`
	var book entities.Book
	err := Instance.QueryRow(query, id).Scan(&book.Id, &book.Title, &book.Isnb)
	if err != nil {
		log.Fatal(err)
	}
	return book
}
