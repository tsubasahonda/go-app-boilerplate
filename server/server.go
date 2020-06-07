package server

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tsubasahonda/go-app-boilerplate/db"
	"github.com/tsubasahonda/go-app-boilerplate/server/handlers"
)

// Init Server
func Init(dbConn *sql.DB) error {
	r := NewRouter()
	d := db.NewSQLDataStorage(dbConn)

	attachHandlers(r, d)

	s := &http.Server{
		Handler: r,
	}

	return s.ListenAndServe()
}

func attachHandlers(r *mux.Router, d db.DataStorage) {
	r.HandleFunc("/news", handlers.ServeNews(d))
}
