// Command menu-creator provides CLI tools for converting PDF menus to Go seed files
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/grow-assistant/menu-creator/internal/pdfparser"
	"github.com/grow-assistant/menu-creator/internal/seeding"
)

func main() {
	log.Println("Menu Creator - PDF to Go Seed File Generator (OpenAI-powered)")
	
	if len(os.Args) < 2 {
		log.Fatal("Usage: menu-creator <pdf-file>")
	}

	pdfPath := os.Args[1]
	if !filepath.IsAbs(pdfPath) {
		absPath, err := filepath.Abs(pdfPath)
		if err != nil {
			log.Fatalf("Failed to get absolute path: %v", err)
		}
		pdfPath = absPath
	}

	log.Printf("Parsing PDF file: %s\n", pdfPath)
	log.Println("Using OpenAI for menu extraction...")
	location, err := pdfparser.ParsePDFMenuWithOpenAI(pdfPath)
	if err != nil {
		log.Printf("OpenAI parsing failed, falling back to direct parsing: %v\n", err)
		location, err = pdfparser.ParsePDFMenu(pdfPath)
		if err != nil {
			log.Fatalf("Failed to parse PDF: %v", err)
		}
	}

	// Pretty print the parsed menu data
	data, err := json.MarshalIndent(location, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal menu data: %v", err)
	}
	fmt.Printf("\nParsed Menu Data:\n%s\n", string(data))

	// Print summary
	fmt.Printf("\nSummary:\n")
	fmt.Printf("Location: %s (%s)\n", location.Name, location.Address)
	for _, menu := range location.Menus {
		fmt.Printf("\nMenu: %s\n", menu.Name)
		for _, category := range menu.Categories {
			fmt.Printf("  Category: %s (%d items)\n", category.Name, len(category.Items))
			for _, item := range category.Items {
				fmt.Printf("    - %s ($%.2f)\n", item.Name, item.Price)
				for _, option := range item.Options {
					fmt.Printf("      * %s (min: %d, max: %d)\n", 
						option.Name, option.MinSelect, option.MaxSelect)
				}
			}
		}
	}

	// Generate Go seed file
	outputDir := filepath.Join("internal", "seeds")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	baseName := strings.TrimSuffix(filepath.Base(pdfPath), ".pdf")
	outputPath := filepath.Join(outputDir, fmt.Sprintf("%s.go", strings.ToLower(baseName)))
	
	// Generate Go seed file
	if err := seeding.ExportToSeedFile(location, pdfPath, outputPath); err != nil {
		log.Fatalf("Failed to generate seed file: %v", err)
	}
	log.Printf("Successfully generated seed file: %s\n", outputPath)
}
