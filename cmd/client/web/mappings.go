//go:build js && wasm

package web

import (
	"net/http"
	"strings"
	"syscall/js"
)

func NewError(msg string) js.Value {
	return Error().New(msg)
}

func NewPromise(handler js.Func) js.Value {
	return Promise().New(handler)
}

func NewResponseInit(resp *http.Response) js.Value {
	headerMap := make(map[string]any, len(resp.Header))
	for key, value := range resp.Header {
		headerMap[key] = strings.Join(value, ",")
	}

	headers := Headers().New(js.ValueOf(headerMap))
	return js.ValueOf(map[string]any{
		"status":     resp.StatusCode,
		"statusText": resp.Status,
		"headers":    headers,
	})
}
