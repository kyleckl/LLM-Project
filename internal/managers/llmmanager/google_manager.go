package llmmanager

import (
	"context"
	"llmApp/internal/errors"

	"google.golang.org/genai"
)

type GoogleAIManager struct {
	model  string
	client *genai.Client
}

type GooglePromptAdapter struct {
	prompt []*genai.Content
	model  string
}

func NewGoogleAIManager(model string, apiKey string) (*GoogleAIManager, error) {
	ctx := context.Background()
	cc := &genai.ClientConfig{
		APIKey: apiKey,
	}

	client, err := genai.NewClient(ctx, cc)
	if err != nil {
		return nil, err
	}

	return &GoogleAIManager{
		model:  model,
		client: client,
	}, nil
}

func (g *GoogleAIManager) ConstructPrompt(userInput string) any {
	prompt := genai.Text(userInput)
	return GooglePromptAdapter{
		prompt: prompt,
		model:  g.model,
	}
}

func (g *GoogleAIManager) SendQuery(ctx context.Context, prompt any) (string, error) {
	googlePrompt, ok := prompt.(GooglePromptAdapter)
	if !ok {
		return "", errors.ErrTypeAssertionFailed
	}

	resp, err := g.client.Models.GenerateContent(ctx, googlePrompt.model, googlePrompt.prompt, nil)
	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
