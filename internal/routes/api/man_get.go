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

		// TODO: Validate arg
		if arg == "" {
			internal.JSONErr(w, http.StatusBadRequest, "query param 'arg' is not defined")
			return
		}

		errMsg := "could not generate man page"

		dto, err := services.GenerateManPage(arg)
		if err != nil {
			slog.Error(errMsg, "err", err)
			internal.JSONErr(w, http.StatusInternalServerError, errMsg)
			return
		}

		if err := internal.JSONDat(w, http.StatusOK, dto); err != nil {
			slog.Error(errMsg, "err", err)
			internal.JSONErr(w, http.StatusInternalServerError, errMsg)
			return
		}
	})
}
