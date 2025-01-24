package pdfparser

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	
	"golang.org/x/exp/constraints"

	"github.com/grow-assistant/menu-creator/internal/config"
	"github.com/grow-assistant/menu-creator/internal/models"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/sashabaranov/go-openai"
)

// Package-level variables removed in favor of dynamic extraction

// ExtractPDFText extracts raw text content from a PDF file
func ExtractPDFText(pdfPath string) (string, error) {
	if !filepath.IsAbs(pdfPath) {
		return "", fmt.Errorf("PDF path must be absolute: %s", pdfPath)
	}

	log.Printf("Extracting text from %s\n", pdfPath)

	// Create temporary directory for content extraction
	tempDir, err := os.MkdirTemp("", "pdfcpu_extract_*")
	if err != nil {
		return "", fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Extract text using pdfcpu
	conf := model.NewDefaultConfiguration()
	// Extract content from PDF using pdfcpu API
	if err := api.ExtractContentFile(pdfPath, tempDir, []string{"1-"}, conf); err != nil {
		return "", fmt.Errorf("failed to extract text from PDF: %w", err)
	}

	// Find and combine extracted content files
	files, err := os.ReadDir(tempDir)
	if err != nil {
		return "", fmt.Errorf("failed to read temp directory: %w", err)
	}

	var content strings.Builder
	for _, file := range files {
		if strings.Contains(file.Name(), "_Content_page_") {
			data, err := os.ReadFile(filepath.Join(tempDir, file.Name()))
			if err != nil {
				return "", fmt.Errorf("failed to read extracted content: %w", err)
			}
			content.Write(data)
			content.WriteString("\n")
		}
	}

	result := content.String()
	if result == "" {
		return "", fmt.Errorf("no text content extracted from PDF")
	}
	
	log.Printf("Successfully extracted text from PDF\n")
	log.Printf("Extracted text sample (first 500 chars):\n%s\n", result[:min(500, len(result))])
	return result, nil
}

// ParsePDFMenuWithOpenAI uses OpenAI to extract structured menu data from PDF text
func ParsePDFMenuWithOpenAI(pdfPath string) (*models.Location, error) {
	// Extract text from PDF
	extractedText, err := ExtractPDFText(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("failed to extract text from PDF: %w", err)
	}

	// Get OpenAI configuration
	openAIConfig, err := config.GetOpenAIConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get OpenAI config: %w", err)
	}

	// Create OpenAI client
	client := openai.NewClient(openAIConfig.APIKey)

	// Construct the system prompt
	systemPrompt := `You are a menu structure analyzer. Extract menu information from the provided text and format it according to this structure:
{
  "name": "Venue Name",
  "address": "Venue Address",
  "menus": [{
    "name": "Menu Name",
    "categories": [{
      "name": "Category Name",
      "description": "Category Description",
      "items": [{
        "name": "Item Name",
        "description": "Item Description",
        "price": 0.00,
        "options": [{
          "name": "Option Group Name",
          "description": "Option Description",
          "minSelect": 0,
          "maxSelect": 1,
          "items": [{
            "name": "Option Item Name",
            "description": "Option Description",
            "price": 0.00
          }]
        }]
      }]
    }]
  }]
}

Important rules and guidelines:

1. Menu Structure:
   - Normalize category names (e.g., "STARTERS" -> "Starters", "ENTREES" -> "Main Courses")
   - Clean and format descriptions properly
   - Extract prices as numbers (e.g., "$12.99" -> 12.99)

2. Dynamic Option Detection:
   - Identify customization options mentioned in item descriptions
   - Detect flavor variants and choices (e.g., if a drink comes in different flavors)
   - Look for preparation preferences (e.g., cooking temperatures, preparation styles)
   - Find common modifications (e.g., protein choices, dressings, sides)

3. Option Group Creation:
   - Create appropriate option groups based on item context
   - Set minSelect=1 for required choices (e.g., side dishes that must be chosen)
   - Set minSelect=0 for optional modifications
   - Use maxSelect=1 for exclusive choices (pick one)
   - Use maxSelect=0 for multiple selections allowed

4. Smart Inference:
   - If an item mentions "choice of", create an option group
   - When descriptions list variants, create flavor/style options
   - For items with common customizations (proteins, sides, toppings, etc.), look for relevant options
   - Identify "add-on" or "extra" items and price them accordingly

5. Return only valid JSON, no additional text

Example option patterns to look for:
- "Choice of..." or "Select..." phrases
- Lists of flavors or varieties
- Preparation instructions ("cooked to order", "temperature")
- Common modifications ("add", "substitute", "extra")
- Dietary options ("gluten-free", "vegetarian")`

	// Create the chat completion request
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openAIConfig.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role: openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Extract menu information from this text and return it as JSON following the specified structure:\n\n%s",
						extractedText),
				},
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("OpenAI API call failed: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no response from OpenAI")
	}

	// Parse the JSON response
	var location models.Location
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &location); err != nil {
		return nil, fmt.Errorf("failed to parse OpenAI response: %w\nResponse: %s", err, resp.Choices[0].Message.Content)
	}

	// Generate IDs for the location and its components
	location.ID = models.GenerateID(location.Name)
	for i := range location.Menus {
		menu := &location.Menus[i]
		menu.ID = models.GenerateID(menu.Name)
		menu.LocationID = location.ID

		for j := range menu.Categories {
			category := &menu.Categories[j]
			category.ID = models.GenerateID(category.Name)
			category.MenuID = menu.ID

			for k := range category.Items {
				item := &category.Items[k]
				item.ID = models.GenerateID(item.Name)
				item.CategoryID = category.ID

				for l := range item.Options {
					option := &item.Options[l]
					option.ID = models.GenerateID(option.Name)
					option.ItemID = item.ID
				}
			}
		}
	}

	return &location, nil
}

// ParsePDFMenu extracts menu data from a PDF file's path
func ParsePDFMenu(pdfPath string) (*models.Location, error) {
	// Extract text from PDF
	extractedText, err := ExtractPDFText(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("failed to extract text from PDF: %w", err)
	}

	// Create location from filename
	locationName := strings.TrimSuffix(filepath.Base(pdfPath), ".pdf")
	location := &models.Location{
		ID:      generateID(locationName),
		Name:    locationName,
		Address: "", // TODO: Extract from PDF if available
		Menus:   []models.Menu{},
	}

	// Create default menu
	menu := models.Menu{
		ID:         generateID(locationName + "Menu"),
		Name:       locationName + " Menu",
		LocationID: location.ID,
		Categories: []models.Category{},
	}

	// Split text into lines and process
	lines := strings.Split(extractedText, "\n")
	var currentCategory *models.Category
	var currentItem *models.MenuItem
	var description strings.Builder
	
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		// Check if line is a category header (all caps)
		if isCategory(line) {
			if currentCategory != nil {
				menu.Categories = append(menu.Categories, *currentCategory)
			}
			currentCategory = &models.Category{
				ID:          generateID(line),
				Name:        normalizeCategory(line),
				Description: "",
				MenuID:      menu.ID,
				Items:       []models.MenuItem{},
			}
			continue
		}

		// Try to extract price from line
		price, itemName := extractPrice(line)
		if price > 0 {
			// If we were building a description, finish it
			if description.Len() > 0 && currentItem != nil {
				currentItem.Description = strings.TrimSpace(description.String())
				description.Reset()
			}

			// This line contains an item with price
			currentItem = &models.MenuItem{
				ID:          generateID(itemName),
				Name:        itemName,
				Description: "",
				Price:       price,
				CategoryID:  currentCategory.ID,
				Options:     []models.Option{},
			}

			if currentCategory != nil {
				currentCategory.Items = append(currentCategory.Items, *currentItem)
			}
		} else if currentItem != nil && !isCategory(line) {
			// Append to multi-line description
			if description.Len() > 0 {
				description.WriteString(" ")
			}
			description.WriteString(line)
		}
	}

	// Handle final description and category
	if description.Len() > 0 && currentItem != nil {
		currentItem.Description = strings.TrimSpace(description.String())
	}
	if currentCategory != nil {
		menu.Categories = append(menu.Categories, *currentCategory)
	}

	location.Menus = append(location.Menus, menu)
	return location, nil
}

// Helper functions

func normalizeCategory(category string) string {
	return models.NormalizeCategory(category)
}

func generateID(name string) string {
	return models.GenerateID(name)
}

func isCategory(line string) bool {
	// Categories are typically in all caps, don't contain prices, and are short
	line = strings.TrimSpace(line)
	return strings.ToUpper(line) == line && 
		!containsPrice(line) && 
		len(line) > 0 && 
		len(line) < 30 && // Categories are usually short
		!strings.Contains(line, "  ") && // Categories don't have multiple spaces
		!strings.Contains(strings.ToLower(line), "page") // Exclude page numbers
}

func containsPrice(line string) bool {
	// Check if line contains a number that could be a price
	reg := regexp.MustCompile(`\d+(\.\d{2})?`)
	return reg.MatchString(line)
}

// min returns the smaller of x or y
func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func extractPrice(line string) (float64, string) {
	// Find price pattern (e.g., "$12.99" or "12.99" or "$12")
	reg := regexp.MustCompile(`\$?(\d+(?:\.\d{2})?)\s*$`)
	matches := reg.FindStringSubmatch(line)
	if len(matches) > 1 {
		price, err := strconv.ParseFloat(matches[1], 64)
		if err == nil && price > 0 && price < 1000 { // Sanity check on price range
			// Remove price and any trailing spaces or dots from description
			description := strings.TrimRight(reg.ReplaceAllString(line, ""), " .")
			description = strings.TrimSpace(description)
			return price, description
		}
	}
	return 0, strings.TrimSpace(line)
}
