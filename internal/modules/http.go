package modules

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// HTTPClient is a wrapper around http.Client for making requests.
type HTTPClient struct {
	client  *http.Client
	BaseURL string
}

// NewHTTPClient creates a new HTTPClient.
func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		client:  &http.Client{},
		BaseURL: baseURL,
	}
}

// Post sends a POST request to a given path with a JSON body.
func (c *HTTPClient) Post(ctx context.Context, path string, body any) (*http.Response, error) {
	bodyReader, err := StructToReader(body)
	if err != nil {
		return nil, err
	}

	url := c.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return c.client.Do(req)
}

// Get sends a GET request to a given path.
func (c *HTTPClient) Get(ctx context.Context, path string, body any) (*http.Response, error) {
	bodyReader, err := StructToReader(body)
	if err != nil {
		return nil, err
	}

	url := c.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return c.client.Do(req)
}

// StructToReader converts a struct to a bytes.Reader containing its JSON representation.
func StructToReader(v any) (*bytes.Reader, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal struct: %w", err)
	}
	return bytes.NewReader(jsonBytes), nil
}
