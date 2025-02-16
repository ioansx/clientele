//go:build js && wasm

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"syscall/js"
	// "github.com/ioansx/clientele/internal/models"
)

var httpClient = http.Client{}

func main() {
	js.Global().Set("manGet", ManGet())

	fmt.Println("Clientele is ready to be served.")

	select {}
}

func ManGet() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {

		// validation
		// if len(args) != 1 {
		// 	errorObj := js.Global().Get("Error").New("error too few args")
		// 	return js.Global().Get("Promise").Get("reject").Invoke(errorObj)
		// }
		//
		// if args[0].IsUndefined() {
		// 	errorObj := js.Global().Get("Error").New("first arg is undefined")
		// 	return js.Global().Get("Promise").Get("reject").Invoke(errorObj)
		// }

		arg := args[0].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) any {
			resolve := args[0]
			reject := args[1]

			go func() {
				url := fmt.Sprintf("%s?%s", "/api/v1/man", url.Values{"arg": {arg}}.Encode())
				req, err := http.NewRequest(http.MethodGet, url, nil)
				if err != nil {
					errorObj := js.Global().Get("Error").New(fmt.Errorf("New request: %w", err).Error())
					reject.Invoke(errorObj)
					return
				}

				resp, err := httpClient.Do(req)
				if err != nil {
					errorObj := js.Global().Get("Error").New(fmt.Errorf("Do request: %w", err).Error())
					reject.Invoke(errorObj)
					return
				}

				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					errorObj := js.Global().Get("Error").New(fmt.Errorf("Read body: %w", err).Error())
					reject.Invoke(errorObj)
					return
				}

				dataJS := js.Global().Get("Uint8Array").New(len(body))
				js.CopyBytesToJS(dataJS, body)

				response := js.Global().Get("Response").New(dataJS, map[string]any{
					"status":     resp.StatusCode,
					"statusText": resp.Status,
					// "headers":    js.Global().Get("Headers").New(resp.Header),
				})

				resolve.Invoke(response)
			}()

			return nil
		})

		return js.Global().Get("Promise").New(handler)
	})
}
