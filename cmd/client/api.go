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
			wrapped := clientele.WrapError("'page' is undefined")
			return web.PromiseReject(wrapped)
		}

		page := args[0].String()

		err := validations.ValidateManGet(page)
		if err != nil {
			wrapped := clientele.WrapError(err.Error())
			return web.PromiseReject(wrapped)
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
