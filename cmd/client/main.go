//go:build js && wasm

package main

import (
	"fmt"

	"github.com/ioansx/clientele/internal/models"
)

func main() {
	// wait := make(chan struct{}, 0)

	// js.Global().Set("ManGet", js.FuncOf(ManGet))

	fmt.Println("Loaded clientele.")

	// <-wait
}

func ManGet() any {
	dat := models.ManGetOutdto{Output: "hey there"}
	return models.Outdto[models.ManGetOutdto]{Dat: dat}
}
