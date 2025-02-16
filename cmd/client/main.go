//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/ioansx/clientele/cmd/client/clientele"
)

func main() {
	clt := clientele.NewClient()
	clt.Handle("manGet", manGet(clt))

	js.Global().Set("clientele", clt.Endpoints())
	fmt.Println("Clientele is ready to be served.")

	select {}
}
