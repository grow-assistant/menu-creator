package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grow-assistant/menu-creator/internal/config"
)

func main() {
	log.Println("Verifying OpenAI Configuration...")
	
	cfg, err := config.GetOpenAIConfig()
	if err != nil {
		log.Fatalf("Failed to get OpenAI config: %v", err)
	}

	if cfg.APIKey == "" {
		log.Fatal("OpenAI API key is not set")
	}

	fmt.Printf("OpenAI Configuration:\n")
	fmt.Printf("API Key: %s...[hidden]\n", cfg.APIKey[:8])
	fmt.Printf("Model: %s\n", cfg.Model)
	
	log.Println("✓ OpenAI configuration verified successfully")
	os.Exit(0)
}
