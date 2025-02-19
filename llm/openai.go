package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"statsbuddy/config"

	"github.com/sashabaranov/go-openai"
)

type OpenAIProvider struct {
	client       *openai.Client
	systemPrompt string
}

func NewOpenAIProvider() (*OpenAIProvider, error) {
	systemPrompt, err := getPromptWithSchema()
	if err != nil {
		return nil, fmt.Errorf("failed to load database schema: %v", err)
	}

	apiKey := config.GetConfig().OpenAIAPIKey
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}
	client := openai.NewClient(apiKey)
	return &OpenAIProvider{
		client:       client,
		systemPrompt: systemPrompt,
	}, nil
}

func (p *OpenAIProvider) GenerateMongoQuery(prompt string) (string, error) {
	resp, err := p.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: p.systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}

func (p *OpenAIProvider) FormatResponse(data interface{}, prompt string) (string, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal data: %v", err)
	}

	resp, err := p.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: ResponseFormattingSystemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "UserPrompt: " + string(prompt),
				},
				{
					Role:    openai.ChatMessageRoleTool,
					Content: "Retrieved data: " + string(dataJSON),
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}
