//go:build js && wasm

package main

import (
// "syscall/js"
// "github.com/ioansx/clientele/internal/models"
)

func main() {
	// wait := make(chan struct{}, 0)

	// js.Global().Set("ManGet", js.FuncOf(ManGet))

	// <-wait
}

//export ManGet
func ManGet() int {
	// dat := models.ManGetOutdto{Output: "hey there"}
	// return models.Outdto[models.ManGetOutdto]{Dat: dat}
	return 5
}
