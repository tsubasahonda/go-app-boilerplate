package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tsubasahonda/go-app-boilerplate/db"
	"github.com/tsubasahonda/go-app-boilerplate/logs"
	"github.com/tsubasahonda/go-app-boilerplate/server/handlers"
)

// Init Server
func Init(port int, dbConn *sql.DB) error {
	r := NewRouter()
	d := db.NewSQLDataStorage(dbConn)

	attachHandlers(r, d)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	logs.Info("Server started...")
	logs.Info("Listening on %s", s.Addr)

	return s.ListenAndServe()
}

func attachHandlers(r *mux.Router, d db.DataStorage) {
	r.HandleFunc("/news", handlers.ServeNews(d))
}
