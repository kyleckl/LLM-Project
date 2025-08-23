package llmmanager

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"llmApp/internal/managers/httpManager"
	"net/http"
)

type LlamaManager struct {
	httpClient *httpManager.HTTPClient
	Model      string
}

type LlamaPromptAdapter struct {
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
	Stream bool   `json:"stream"`
}

type LlamaResponseAdapter struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

func NewLlamaClient(model string) *LlamaManager {
	// The base URL for Ollama is localhost:11434
	// Make sure the Ollama server is running locally
	baseURL := "http://localhost:11434"

	return &LlamaManager{
		httpClient: httpManager.NewHTTPClient(baseURL),
		Model:      model,
	}
}

func (l *LlamaManager) ConstructPrompt(userInput string, stream bool) any {
	return LlamaPromptAdapter{
		Prompt: userInput,
		Model:  l.Model,
		Stream: stream,
	}
}

func (l *LlamaManager) SendQuery(ctx context.Context, prompt any) (*http.Response, error) {
	// Depending on the model, the API endpoint may differ
	path := "/api/generate"

	resp, err := l.httpClient.Post(ctx, path, prompt)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Given an HTTP response from the LLM, it streams and returns the response as a string
func (l *LlamaManager) StreamResponse(ctx context.Context, resp *http.Response) (response string, err error) {
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Bytes()
		var chunk LlamaResponseAdapter
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
