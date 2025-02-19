package llm

import (
	"encoding/json"
	"os"
	"path"
	"strings"
	"text/template"
)

// loadDatabaseSchema loads the database schema from the config file
func loadDatabaseSchema() (string, error) {
	// path get current working directory
	wd, _ := os.Getwd()
	path := path.Join(path.Dir(path.Dir(wd)), "config", "mongodb_schema.json")
	schemaBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Verify schema is valid JSON
	var schemaMap map[string]interface{}
	if err := json.Unmarshal(schemaBytes, &schemaMap); err != nil {
		return "", err
	}

	return string(schemaBytes), nil
}

// getPromptWithSchema returns the system prompt with the database schema inserted
func getPromptWithSchema() (string, error) {
	schema, err := loadDatabaseSchema()
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("prompt").Parse(QueryGenerationSystemPrompt)
	if err != nil {
		return "", err
	}

	data := struct {
		DBSchema string
	}{
		DBSchema: schema,
	}

	var result strings.Builder
	if err := tmpl.Execute(&result, data); err != nil {
		return "", err
	}

	return result.String(), nil
}
