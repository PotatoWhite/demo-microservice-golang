package handler

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()

	dbMock := sqlx.NewDb(mockDB, "sqlmock")

	tests := []struct {
		name string
		want string
	}{
		{"1", "{\"status\":200,\"message\":\"Hello visitor! : 1\"}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
			r.Header.Set("User-Agent", "gotest")
			w := httptest.NewRecorder()

			query := `^INSERT .*`
			mock.ExpectQuery(query).
				WithArgs(r.UserAgent()).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			home := Home(dbMock)
			home(w, r)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("errors while running SQL queries:\n %s", err)
			}

			body := w.Body.String()
			if body != tt.want {
				t.Errorf("Home() = %q, want %q", body, tt.want)
			}
		})
	}
}
