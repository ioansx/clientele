//go:build js && wasm

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"syscall/js"
	// "github.com/ioansx/clientele/internal/models"
)

var httpClient = http.Client{}

func main() {
	wait := make(chan int)

	js.Global().Set("manGet", js.FuncOf(ManGet))

	fmt.Println("Clientele is ready to be served.")

	<-wait
}

func ManGet(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return "error too few args"
	}

	if args[0].IsUndefined() {
		return "first arg is undefined"
	}

	c := make(chan errResponse, 1)
	go func() {
		arg := args[0].String()
		data, err := doRequest(http.MethodGet, "/api/v1/man", url.Values{"arg": {arg}})
		if err != nil {
			c <- errResponse{dat: "", err: fmt.Errorf("Client error: %w", err)}
			return
		}
		c <- errResponse{dat: data, err: nil}
	}()
	data := <-c

	return data
}

type errResponse struct {
	dat string
	err error
}

func doRequest(method string, path string, query url.Values) (string, error) {
	url := fmt.Sprintf("%s?%s", path, query.Encode())
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", fmt.Errorf("doRequest NewRequest: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("doRequest Do: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("doRequest read body: %w", err)
	}

	return string(body), nil
}
