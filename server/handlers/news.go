package handlers

import (
	"net/http"

	"github.com/tsubasahonda/go-app-boilerplate/db"
)

// ServeNews is handler of news endpoint
func ServeNews(d db.DataStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
