package llm

// LLMProvider interface defines methods that any LLM provider must implement
type LLMProvider interface {
	GenerateMongoQuery(prompt string) (string, error)
	FormatResponse(data interface{}, prompt string) (string, error)
}

// LLMType represents supported LLM providers
type LLMType string

const (
	OpenAI LLMType = "openai"
	Gemini LLMType = "gemini"
)
