//go:build js && wasm

package web

import "syscall/js"

func Console() js.Value {
	return js.Global().Get("console")
}

func ConsoleLog(args ...any) js.Value {
	return Console().Call("log", args)
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

func PromiseReject(args ...any) js.Value {
	return Promise().Call("reject", args)
}

func Response() js.Value {
	return js.Global().Get("Response")
}

func Uint8Array() js.Value {
	return js.Global().Get("Uint8Array")
}
