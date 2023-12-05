package database

import (
	"database/sql"
	"fmt"
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
	Instance = db
}

func createMigrations(connString string) {
	m, err := migrate.New("file://database/migrations", connString)
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	m.Up()
}
