//go:build js && wasm

package main

import (
	"net/http"
	"net/url"
	"syscall/js"

	"github.com/ioansx/clientele/cmd/client/web"
)

func manGet(this js.Value, args []js.Value) any {
	if len(args) != 1 || args[0].IsUndefined() {
		return web.PromiseReject(web.NewError("'arg' is undefined"))
	}

	arg := args[0].String()

	path := "/api/v1/man"
	query := url.Values{"arg": {arg}}

	handler := newPromiseHandler(http.MethodGet, path, query)

	return web.NewPromise(handler)
}
