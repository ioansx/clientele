//go:build js && wasm

package main

import (
	"fmt"
	"net/http"
	"syscall/js"
)

var httpClient = http.Client{}

func main() {
	client := NewClient()
	client.Handle("manGet", manGet)

	js.Global().Set("clientele", client.Clientele())

	fmt.Println("Clientele is ready to be served.")

	select {}
}
