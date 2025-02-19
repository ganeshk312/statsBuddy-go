package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"statsbuddy/config"

	generativelanguage "cloud.google.com/go/ai/generativelanguage/apiv1"
	"cloud.google.com/go/ai/generativelanguage/apiv1/generativelanguagepb"
	"google.golang.org/api/option"
)

const MongoQuerySystemPromptGemini = "You are a MongoDB query generator. Generate a MongoDB query based on the given natural language query."

type GeminiProvider struct {
	client       *generativelanguage.GenerativeClient
	ctx          context.Context
	systemPrompt string
}

func NewGeminiProvider() (*GeminiProvider, error) {
	apiKey := config.GetConfig().GeminiAPIKey
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY environment variable not set")
	}
	systemPrompt, err := getPromptWithSchema()
	if err != nil {
		return nil, fmt.Errorf("failed to get prompt with schema: %v", err)
	}
	ctx := context.Background()
	client, err := generativelanguage.NewGenerativeClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %v", err)
	}

	return &GeminiProvider{
		client:       client,
		ctx:          ctx,
		systemPrompt: systemPrompt,
	}, nil
}

func (p *GeminiProvider) GenerateMongoQuery(prompt string) (string, error) {
	req := &generativelanguagepb.GenerateContentRequest{
		Model: "models/gemini-pro",
		Contents: []*generativelanguagepb.Content{
			{
				Parts: []*generativelanguagepb.Part{
					{
						Data: &generativelanguagepb.Part_Text{
							Text: fmt.Sprintf("%s\n\nQuery: %s, \nCode:", p.systemPrompt, prompt),
						},
					},
				},
			},
		},
	}

	resp, err := p.client.GenerateContent(p.ctx, req)
	if err != nil {
		return "", fmt.Errorf("gemini API error: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response from Gemini")
	}

	part := resp.Candidates[0].Content.Parts[0]
	if textPart, ok := part.Data.(*generativelanguagepb.Part_Text); !ok || textPart.Text == "" {
		return "", fmt.Errorf("no text in Gemini response")
	} else {
		return textPart.Text, nil
	}
}

func (p *GeminiProvider) FormatResponse(data interface{}, prompt string) (string, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal data: %v", err)
	}

	req := &generativelanguagepb.GenerateContentRequest{
		Model: "models/gemini-pro",
		Contents: []*generativelanguagepb.Content{
			{
				Parts: []*generativelanguagepb.Part{
					{
						Data: &generativelanguagepb.Part_Text{
							Text: fmt.Sprintf("%s\n\nUserPrompt: %s\n\nRetrieved Data: %s", ResponseFormattingSystemPrompt, prompt, string(dataJSON)),
						},
					},
				},
			},
		},
	}

	resp, err := p.client.GenerateContent(p.ctx, req)
	if err != nil {
		return "", fmt.Errorf("gemini API error: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response from Gemini")
	}

	part := resp.Candidates[0].Content.Parts[0]
	if textPart, ok := part.Data.(*generativelanguagepb.Part_Text); !ok || textPart.Text == "" {
		return "", fmt.Errorf("no text in Gemini response")
	} else {
		return textPart.Text, nil
	}
}
