package config

import (
	"os"
)

type Config struct {
	OpenAIAPIKey string
	GeminiAPIKey string
	MongoDBURI   string
	DatabaseName string
	LLMType      string
}

func LoadConfig() *Config {
	return &Config{
		OpenAIAPIKey: os.Getenv("OPENAI_API_KEY"),
		// GeminiAPIKey: os.Getenv("GEMINI_API_KEY"),
		GeminiAPIKey: "AIzaSyC4kBx9wp4ySHiQ-b3u0Dq7jjJ3GQs6HfU",
		// MongoDBURI:   os.Getenv("MONGODB_URI"),
		MongoDBURI: "mongodb://localhost:27017",
		// DatabaseName: os.Getenv("MONGODB_DATABASE"),
		DatabaseName: "all_india",
		// LLMType:     os.Getenv("LLM_TYPE"),
		LLMType: "gemini",
	}
}

var config *Config

func init() {
	config = LoadConfig()
}

func GetConfig() *Config {
	return config
}
