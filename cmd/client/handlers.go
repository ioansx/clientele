//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/ioansx/clientele/cmd/client/clientele"
	"github.com/ioansx/clientele/cmd/client/web"
)

func newPromiseHandler(client *clientele.Client, req clientele.Request) js.Func {
	callAPI := func(resolve, reject js.Value) {
		resp, err := client.Do(req)
		if err != nil {
			errorObj := web.NewError(fmt.Errorf("Clientele: %w", err).Error())
			reject.Invoke(errorObj)
			return
		}

		response := web.Response().New(resp.Body, web.NewResponseInit(resp))
		resolve.Invoke(response)
	}

	return js.FuncOf(func(this js.Value, args []js.Value) any {
		resolve := args[0]
		reject := args[1]
		go callAPI(resolve, reject)
		return nil
	})
}
