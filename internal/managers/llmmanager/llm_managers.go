package llmmanager

import (
	"context"
	"net/http"
)

type LLMManager interface {
	ConstructPrompt(userInput string, stream bool) any
	SendQuery(ctx context.Context, prompt any) (*http.Response, error)
	StreamResponse(ctx context.Context, resp *http.Response) (response string, err error)
}
