package internal

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ioansx/clientele/internal/models"
)

func JSONDat(w http.ResponseWriter, code int, dat any) error {
	return JSONOutdto(w, code, models.Outdto[any]{Dat: dat})
}

func JSONErr(w http.ResponseWriter, code int, errMsg string) {
	err := JSONOutdto(w, code, models.Outdto[any]{Err: []string{errMsg}})
	if err != nil {
		slog.Error("serialization failed", "err", err)
	}
}

func JSONOutdto(w http.ResponseWriter, code int, outdto any) error {
	marshaled, err := json.Marshal(outdto)
	if err != nil {
		errMsg := "serialization failed"
		slog.Error(errMsg, "err", err)
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	h := w.Header()
	h.Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintln(w, string(marshaled))

	return nil
}

func TextBadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "500 bad request", http.StatusBadRequest)
}

func TextInternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "500 internal server error", http.StatusInternalServerError)
}
