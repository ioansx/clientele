package routes

import (
	"html/template"
	"net/http"

	"github.com/ioansx/clientele/internal"
)

func AddRoutes(mux *http.ServeMux) {
	var templates = template.Must(template.ParseGlob("web/templates/*.tmpl"))

	mux.Handle("GET /{$}", internal.StandardMiddlewares(indexHandler(templates)))
	mux.Handle("GET /man", internal.StandardMiddlewares(manPageHandler(templates)))
	mux.Handle("GET /", http.FileServer(http.Dir("web/static")))
}
