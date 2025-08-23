package llmmanager

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"llmApp/internal/managers/httpManager"
	"net/http"
)

type llmManager struct {
	httpClient *httpManager.HTTPClient
	APIKey     string
	Model      string
}

type QueryAdapter struct {
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
	Stream bool   `json:"stream"`
}

type ResponseAdapter struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

func NewLLMClient(model string, apiKey string) *llmManager {
	// Change the HTTP client base URL based on the model specified
	// Currently only "ollamaModel" is supported
	var baseURL string
	switch model {
	case "ollamaModel":
		baseURL = "http://localhost:11434"
	}

	return &llmManager{
		httpClient: httpManager.NewHTTPClient(baseURL),
		APIKey:     apiKey,
		Model:      model,
	}
}

func (l *llmManager) SendQuery(ctx context.Context, prompt any) (*http.Response, error) {
	// Depending on the model, the API endpoint may differ
	var path string
	switch l.Model {
	case "ollamaModel":
		path = "/api/generate"
	}

	resp, err := l.httpClient.Post(ctx, path, prompt)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Given an HTTP response from the LLM, it streams and returns the response as a string
func (l *llmManager) StreamResponse(ctx context.Context, resp *http.Response) (response string, err error) {
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Bytes()
		var chunk ResponseAdapter
		err := json.Unmarshal(line, &chunk)
		if err != nil {
			return "", err
		}

		response += chunk.Response
		fmt.Print(chunk.Response)

		if chunk.Done {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return response, nil
}
