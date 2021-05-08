package main

import (
	_ "database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"

	""
)

func main() {
	dataSource := "postgres://goland:goland@%s:5432/goland?sslmode=disable"

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf(dataSource, dbHost))
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", handler.Home(db))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

