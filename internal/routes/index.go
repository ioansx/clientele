package routes

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/ioansx/clientele/internal"
)

func indexHandler(templates *template.Template) http.Handler {
	name := "index.tmpl"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, name, nil)
		if err != nil {
			slog.Error(err.Error(), "template", name)
			internal.TextInternalServerError(w, r)
			return
		}
	})
}
