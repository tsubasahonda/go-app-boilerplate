package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tsubasahonda/go-app-boilerplate/db"
	"github.com/tsubasahonda/go-app-boilerplate/logs"
	"github.com/tsubasahonda/go-app-boilerplate/models"
)

// ServeNews is handler of news endpoint
func ServeNews(d db.DataStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		news, err := d.GetNews()
		if err != nil {
			logs.Error("Request: %s, reading news from database: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		bytes, err := json.Marshal(models.NewNews(news))
		if err != nil {
			logs.Error("Request: %s, serializing news: %v", requestSummary(r), err)
			InternalServerError(w, r)
			return
		}

		w.Header().Set(contentType, jsonContent)

		if _, err = w.Write(bytes); err != nil {
			logs.Error("Request: %s, writing response: %v", requestSummary(r), err)
		}
	}
}
