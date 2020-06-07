package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tsubasahonda/go-app-boilerplate/server/handlers"
)

// Init Server
func Init() error {
	r := NewRouter()

	attachHandlers(r)

	s := &http.Server{
		Handler: r,
	}

	return s.ListenAndServe()
}

func attachHandlers(r *mux.Router) {
	r.HandleFunc("/books", handlers.ServeBooks())
}
