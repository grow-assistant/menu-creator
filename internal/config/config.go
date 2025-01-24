package config

import (
	"sync"
)

var (
	openAIConfig     *OpenAIConfig
	openAIConfigOnce sync.Once
)

// GetOpenAIConfig returns the OpenAI configuration singleton
func GetOpenAIConfig() (*OpenAIConfig, error) {
	var err error
	openAIConfigOnce.Do(func() {
		openAIConfig = DefaultOpenAIConfig()
		err = openAIConfig.Validate()
	})
	if err != nil {
		return nil, err
	}
	return openAIConfig, nil
}
