package seeding

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/grow-assistant/menu-creator/internal/models"
)

func TestExportToSeedFile(t *testing.T) {
	// Create test location with sample data
	location := &models.Location{
		ID:      "testLocation",
		Name:    "Test Location",
		Address: "123 Test St",
		Menus: []models.Menu{
			{
				ID:         "testMenu",
				Name:       "Test Menu",
				LocationID: "testLocation",
				Categories: []models.Category{
					{
						ID:          "testCategory",
						Name:        "Test Category",
						Description: "Test Description",
						MenuID:      "testMenu",
						Items: []models.MenuItem{
							{
								ID:          "itemB",
								Name:        "B Item",
								Description: "Description B",
								Price:       15.99,
								CategoryID:  "testCategory",
								Options:     []models.Option{},
							},
							{
								ID:          "itemA",
								Name:        "A Item",
								Description: "Description A",
								Price:       12.99,
								CategoryID:  "testCategory",
								Options: []models.Option{
									{
										ID:          "optionTest",
										Name:        "Test Option",
										Description: "Option Description",
										MinSelect:   0,
										MaxSelect:   1,
										ItemID:      "itemA",
										Items: []models.OptionItem{
											{
												Name:        "Option B",
												Description: "Option B Desc",
												Price:       2.00,
											},
											{
												Name:        "Option A",
												Description: "Option A Desc",
												Price:       1.00,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Create temporary directory for test output
	tmpDir, err := os.MkdirTemp("", "generator_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	outputPath := filepath.Join(tmpDir, "test_location.go")
	err = ExportToSeedFile(location, "test.pdf", outputPath)
	if err != nil {
		t.Fatalf("ExportToSeedFile failed: %v", err)
	}

	// Read generated file
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Verify file content
	generatedCode := string(content)

	// Check required elements
	requiredElements := []string{
		"// Code generated by menu-creator. DO NOT EDIT.",
		"// Source: test.pdf",
		"// Generated:",
		"package seeding",
		"func CreateTestLocation() *models.Location",
		"A Item",
		"B Item",
		"Option A",
		"Option B",
	}

	for _, required := range requiredElements {
		if !strings.Contains(generatedCode, required) {
			t.Errorf("Generated code missing required element: %s", required)
		}
	}

	// Verify items are sorted (A Item should appear before B Item)
	aIndex := strings.Index(generatedCode, "A Item")
	bIndex := strings.Index(generatedCode, "B Item")
	if aIndex == -1 || bIndex == -1 {
		t.Error("Missing menu items in generated code")
	} else if aIndex > bIndex {
		t.Error("Menu items not properly sorted")
	}

	// Verify options are sorted (Option A should appear before Option B)
	optionAIndex := strings.Index(generatedCode, "Option A")
	optionBIndex := strings.Index(generatedCode, "Option B")
	if optionAIndex == -1 || optionBIndex == -1 {
		t.Error("Missing options in generated code")
	} else if optionAIndex > optionBIndex {
		t.Error("Options not properly sorted")
	}
}
