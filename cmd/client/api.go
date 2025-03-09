//go:build js && wasm

package main

import (
	"net/http"
	"net/url"
	"syscall/js"

	"github.com/ioansx/clientele/cmd/client/clientele"
	"github.com/ioansx/clientele/cmd/client/web"
	"github.com/ioansx/clientele/internal/validations"
)

func manGet(client *clientele.Client) func(js.Value, []js.Value) any {
	return func(this js.Value, args []js.Value) any {
		if len(args) != 1 || args[0].IsUndefined() {
			return web.PromiseReject(web.NewError("'page' is undefined"))
		}

		page := args[0].String()

		err := validations.ValidateManGet(page)
		if err != nil {
			return web.PromiseReject(web.NewError(err.Error()))
		}

		req := clientele.Request{
			Method: http.MethodGet,
			Path:   "/api/v1/man",
			Query:  url.Values{"page": {page}},
		}
		handler := newPromiseHandler(client, req)

		return web.NewPromise(handler)
	}
}
