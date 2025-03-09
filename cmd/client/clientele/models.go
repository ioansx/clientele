//go:build js && wasm

package clientele

import (
	"encoding/json"

	"github.com/ioansx/clientele/internal/models"
)

func WrapError(msg string) string {
	outdto := models.Outdto[any]{Dat: nil, Err: msg}
	marshaled, _ := json.Marshal(outdto)
	return string(marshaled)
}
