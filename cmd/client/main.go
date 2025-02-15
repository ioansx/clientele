//go:build js && wasm

package main

import (
	"syscall/js"
	// "github.com/ioansx/clientele/internal/models"
)

func main() {
	wait := make(chan int, 0)

	// js.Global().Set("ManGet", js.FuncOf(ManGet))

	<-wait
}

func ManGet(this js.Value, args []js.Value) any {
	// dat := models.ManGetOutdto{Output: "hey there"}
	// return models.Outdto[models.ManGetOutdto]{Dat: dat}
	return 5
}
