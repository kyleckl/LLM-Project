package llmmanager

import (
	"context"

	"google.golang.org/genai"
)

type GoogleAIManager struct {
	model   string
	client  *genai.Client
	content []*genai.Content
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

func (g *GoogleAIManager) ConstructPrompt(userInput string) error {
	prompt := genai.Text(userInput)
	g.content = prompt
	return nil
}

func (g *GoogleAIManager) SendQuery(ctx context.Context) (string, error) {
	resp, err := g.client.Models.GenerateContent(ctx, g.model, g.content, nil)
	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
