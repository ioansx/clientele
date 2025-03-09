package api

import (
	"log/slog"
	"net/http"

	"github.com/ioansx/clientele/internal"
	"github.com/ioansx/clientele/internal/services"
	"github.com/ioansx/clientele/internal/validations"
)

func ManGetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var page string
		for key, value := range r.URL.Query() {
			if key == "page" && len(value) == 1 {
				page = value[0]
			}
		}

		errMsg := "could not generate man page"

		err := validations.ValidateManGet(page)
		if err != nil {
			slog.Error(errMsg, "err", err)
			internal.JSONErr(w, http.StatusBadRequest, err.Error())
			return
		}

		dto, err := services.GenerateManPage(page)
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
