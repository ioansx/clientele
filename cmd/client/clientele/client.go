//go:build js && wasm

package clientele

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"syscall/js"
)

type Request struct {
	Method string
	Path   string
	Query  url.Values
	Body   string
}

type Response struct {
	StatusCode int
	Status     string
	Headers    map[string][]string
	Body       string
}

type Client struct {
	httpClient http.Client
	endpoints  map[string]any
}

func NewClient() *Client {
	return &Client{
		httpClient: http.Client{},
		endpoints:  make(map[string]any),
	}
}

func (c *Client) Endpoints() map[string]any {
	return c.endpoints
}

func (c *Client) Handle(name string, fn func(this js.Value, args []js.Value) any) {
	c.endpoints[name] = js.FuncOf(fn)
}

func (c *Client) Do(req Request) (*Response, error) {
	url := req.Path
	if len(req.Query) > 0 {
		url = fmt.Sprintf("%s?%s", req.Path, req.Query.Encode())
	}

	var bodyReader io.Reader
	if req.Body != "" {
		bodyReader = strings.NewReader(req.Body)
	}

	httpReq, err := http.NewRequest(req.Method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("New request: %w", err)
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("Do request: %w", err)
	}

	defer httpResp.Body.Close()
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body: %w", err)
	}

	resp := Response{
		StatusCode: httpResp.StatusCode,
		Status:     httpResp.Status,
		Headers:    httpResp.Header,
		Body:       string(body),
	}

	return &resp, nil
}
