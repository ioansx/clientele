//go:build js && wasm

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"syscall/js"
	// "github.com/ioansx/clientele/internal/models"
	"github.com/ioansx/clientele/cmd/client/web"
)

var httpClient = http.Client{}

func main() {
	js.Global().Set("manGet", js.FuncOf(ManGet))

	fmt.Println("Clientele is ready to be served.")

	select {}
}

func ManGet(this js.Value, args []js.Value) any {
	if len(args) != 1 || args[0].IsUndefined() {
		errorObj := web.Error().New("'arg' is undefined.")
		return web.PromiseReject().Invoke(errorObj)
	}

	arg := args[0].String()

	path := "/api/v1/man"
	query := url.Values{"arg": {arg}}

	handler := makePromiseHandler(http.MethodGet, path, query)

	return web.Promise().New(handler)
}

func makePromiseHandler(method string, path string, query url.Values) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		resolve := args[0]
		reject := args[1]

		go func() {
			url := path
			if len(query) > 0 {
				url = fmt.Sprintf("%s?%s", path, query.Encode())
			}

			req, err := http.NewRequest(method, url, nil)
			if err != nil {
				errorObj := web.Error().New(fmt.Errorf("New request: %w", err).Error())
				reject.Invoke(errorObj)
				return
			}

			resp, err := httpClient.Do(req)
			if err != nil {
				errorObj := web.Error().New(fmt.Errorf("Do request: %w", err).Error())
				reject.Invoke(errorObj)
				return
			}

			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				errorObj := web.Error().New(fmt.Errorf("Read body: %w", err).Error())
				reject.Invoke(errorObj)
				return
			}

			bodyJS := web.Uint8Array().New(len(body))
			js.CopyBytesToJS(bodyJS, body)

			response := web.Response().New(bodyJS, web.ResponseInit(resp))

			resolve.Invoke(response)
		}()

		return nil
	})
}
