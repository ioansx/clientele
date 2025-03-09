//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/ioansx/clientele/cmd/client/clientele"
)

func main() {
	c := clientele.NewClient()
	c.Handle("manGet", manGet(c))

	js.Global().Set("clientele", c.Endpoints())
	fmt.Println("Clientele is ready to be served.")

	select {}
}
