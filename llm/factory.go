package llm

import "fmt"

// LLMFactory creates instances of LLM providers
type LLMFactory struct{}

// CreateLLM creates and returns an LLM provider based on the specified type
func (f *LLMFactory) CreateLLM(llmType LLMType) (LLMProvider, error) {
    switch llmType {
    case OpenAI:
        return NewOpenAIProvider()
    case Gemini:
        return NewGeminiProvider()
    default:
        return nil, fmt.Errorf("unsupported LLM provider: %s", llmType)
    }
}