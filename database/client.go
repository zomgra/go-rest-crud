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
	err := Instance.QueryRow(query, b.Title, b.Isnbn).Scan(&id)
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
		rows.Scan(&book.Id, &book.Title, &book.AuthorId, &book.Isnbn)
		log.Print(book)
		if book.Id == "" {
			break
		}
		log.Print(books)
		books = append(books, book)
	}
	return books
}
