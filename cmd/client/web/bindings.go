//go:build js && wasm

package web

import "syscall/js"

func ConsoleLog() js.Value {
	return js.Global().Get("console").Get("log")
}

func Error() js.Value {
	return js.Global().Get("Error")
}

func Headers() js.Value {
	return js.Global().Get("Headers")
}

func Promise() js.Value {
	return js.Global().Get("Promise")
}

func PromiseReject() js.Value {
	return js.Global().Get("Promise").Get("reject")
}

func Response() js.Value {
	return js.Global().Get("Response")
}

func Uint8Array() js.Value {
	return js.Global().Get("Uint8Array")
}
