package internal

import (
	"fmt"
	"log/slog"
	"net/http"
)

func StandardMiddlewares(handler http.Handler) http.Handler {
	var h http.Handler
	h = TraceMiddleware(handler)
	return h
}

func TraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info(fmt.Sprintf("%s %s", r.Method, r.URL))
		next.ServeHTTP(w, r)
	})
}
