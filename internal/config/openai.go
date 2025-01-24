package config

import (
	"errors"
	"os"
)

// OpenAIConfig holds configuration for OpenAI API interactions
type OpenAIConfig struct {
	APIKey string
	Model  string
}

// DefaultOpenAIConfig returns a new OpenAIConfig with default settings
func DefaultOpenAIConfig() *OpenAIConfig {
	return &OpenAIConfig{
		APIKey: os.Getenv("OPENAI_API_KEY"),
		Model:  getEnvWithDefault("MODEL_FOR_ANALYSIS", "gpt-4"),
	}
}

// Validate ensures all required configuration is present
func (c *OpenAIConfig) Validate() error {
	if c.APIKey == "" {
		return errors.New("OPENAI_API_KEY environment variable is required")
	}
	if c.Model == "" {
		return errors.New("OpenAI model must be specified")
	}
	return nil
}

// getEnvWithDefault returns environment variable value or default if not set
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
