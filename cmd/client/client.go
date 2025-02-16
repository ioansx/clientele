//go:build js && wasm

package main

import "syscall/js"

type Client struct {
	endpoints map[string]any
}

func NewClient() *Client {
	return &Client{endpoints: make(map[string]any)}
}

func (c *Client) Handle(name string, fn func(this js.Value, args []js.Value) any) {
	c.endpoints[name] = js.FuncOf(fn)
}

func (c *Client) Clientele() map[string]any {
	return c.endpoints
}
