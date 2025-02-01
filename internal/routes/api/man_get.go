package api

import (
	"log/slog"
	"net/http"

	"github.com/ioansx/clientele/internal"
	"github.com/ioansx/clientele/internal/services"
)

func ManGetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var arg string
		for key, value := range r.URL.Query() {
			if key == "arg" && len(value) == 1 {
				arg = value[0]
			}
		}

		if arg == "" {
			internal.BadRequest(w, r)
			return
		}

		// TODO: Validate arg

		output, err := services.GenerateManPage(arg)
		if err != nil {
			slog.Error(err.Error())
			internal.InternalServerError(w, r)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		if _, err := w.Write(output); err != nil {
			slog.Error(err.Error())
			internal.InternalServerError(w, r)
			return
		}
	})
}
