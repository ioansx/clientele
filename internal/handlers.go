package internal

import (
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "500 internal server error", http.StatusInternalServerError)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "500 bad request", http.StatusBadRequest)
}
