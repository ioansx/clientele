//go:build js && wasm

package web

import (
	"strings"
	"syscall/js"

	"github.com/ioansx/clientele/cmd/client/clientele"
)

func NewError(msg string) js.Value {
	return Error().New(msg)
}

func NewPromise(handler js.Func) js.Value {
	return Promise().New(handler)
}

func NewResponseInit(resp *clientele.Response) js.Value {
	headerMap := make(map[string]any, len(resp.Headers))
	for key, value := range resp.Headers {
		headerMap[key] = strings.Join(value, ",")
	}

	headers := Headers().New(headerMap)
	return js.ValueOf(map[string]any{
		"status":     resp.StatusCode,
		"statusText": resp.Status,
		"headers":    headers,
	})
}
