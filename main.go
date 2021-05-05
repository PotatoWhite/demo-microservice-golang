package main

import (
	_ "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
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

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		visitorId := 0
		err := db.QueryRow("INSERT INTO visitors(user_agent, datetime) VALUES ($1, now()) RETURNING id",
			request.UserAgent(),
		).Scan(&visitorId)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)

			errorMsg := fmt.Sprintf("Internal error %v", err)
			_, _ = writer.Write([]byte(errorMsg))

			return
		}

		writer.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(writer, fmt.Sprintf("Hello visitor! : %d !", visitorId))
	})



	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
