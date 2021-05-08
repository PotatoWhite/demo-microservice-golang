package handler

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
)

// Home handles the home page
func Home(db *sqlx.DB) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
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
		writer.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(writer, fmt.Sprintf("{\"status\":200,\"message\":\"Hello visitor! : %d\"}", visitorId))
	}
}
