//go:build js && wasm

package main

import (
	"fmt"
	"net/http"
	"syscall/js"
)

var httpClient = http.Client{}

func main() {
	js.Global().Set("manGet", js.FuncOf(manGet))

	fmt.Println("Clientele is ready to be served.")

	select {}
}
