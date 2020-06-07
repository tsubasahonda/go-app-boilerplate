package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tsubasahonda/go-app-boilerplate/logs"
)

type FieldValidationRes struct {
	ErrType string `json:"error_type"`
	Field   string `json:"field"`
}

const (
	validationErrType = "validation"
	contentType       = "Content-Type"
	authToken         = "X-Auth-Token"
	jsonContent       = "application/json"
	mp3Content        = "audio/mpeg"
)

func logRequest(r *http.Request) {
	logs.Info("Received request: %s", requestSummary(r))
}

func requestSummary(r *http.Request) string {
	return fmt.Sprintf("%v %v", r.Method, r.URL)
}

func FieldValidationFailed(w http.ResponseWriter, r *http.Request, field string) error {
	w.Header().Set(contentType, jsonContent)
	w.WriteHeader(http.StatusBadRequest)

	var res = FieldValidationRes{validationErrType, field}
	resJson, err := json.Marshal(res)
	if err != nil {
		return err
	}

	_, err = w.Write(resJson)
	return err
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "400 bad request", http.StatusBadRequest)
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "401 unauthorized", http.StatusUnauthorized)
}

func Forbidden(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "403 forbidden", http.StatusForbidden)
}

func Conflict(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "409 conflict", http.StatusConflict)
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "500 internal server error", http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Not Found", http.StatusNotFound)
}
