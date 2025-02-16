//go:build js && wasm

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"syscall/js"

	"github.com/ioansx/clientele/cmd/client/web"
)

func newPromiseHandler(method string, path string, query url.Values) js.Func {
	callAPI := func(resolve, reject js.Value) {
		url := path
		if len(query) > 0 {
			url = fmt.Sprintf("%s?%s", path, query.Encode())
		}

		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			errorObj := web.NewError(fmt.Errorf("New request: %w", err).Error())
			reject.Invoke(errorObj)
			return
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			errorObj := web.NewError(fmt.Errorf("Do request: %w", err).Error())
			reject.Invoke(errorObj)
			return
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorObj := web.NewError(fmt.Errorf("Read body: %w", err).Error())
			reject.Invoke(errorObj)
			return
		}

		bodyJS := web.Uint8Array().New(len(body))
		js.CopyBytesToJS(bodyJS, body)

		response := web.Response().New(bodyJS, web.NewResponseInit(resp))

		resolve.Invoke(response)
	}

	return js.FuncOf(func(this js.Value, args []js.Value) any {
		resolve := args[0]
		reject := args[1]
		go callAPI(resolve, reject)
		return nil
	})
}
