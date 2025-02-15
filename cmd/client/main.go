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

func main() {
	wait := make(chan int, 0)

	js.Global().Set("manGet", js.FuncOf(ManGet))

	fmt.Println("Clientele is ready to be served.")

	<-wait
}

func ManGet(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return "error too few args"
	}

	if args[0].IsUndefined() {
		return "first arg is undefined"
	}

	arg := args[0].String()
	return get("/api/v1/man", url.Values{"arg": {arg}})
}

func get(path string, query url.Values) string {
	resp, err := http.Get(fmt.Sprintf("%s?%s", path, query.Encode()))
	if err != nil {
		fmt.Println("Error: %w", err)
		return "error"
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return string(body)
}
