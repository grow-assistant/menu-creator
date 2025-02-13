# Go Style Guide for Menu Creator

## Overview
This style guide establishes consistent patterns for .go seed files in the menu-creator project. Following these guidelines ensures maintainability and readability across the codebase.

## Naming Conventions
- Use PascalCase for exported types, interfaces, and functions (e.g., `MenuItem`, `ParsePDFMenu`)
- Use camelCase for unexported (private) functions and variables (e.g., `parsePDF`, `menuItems`)
- Use all uppercase for acronyms (e.g., `PDF`, `JSON`)
- Keep names descriptive and concise
- Prefix interface names with 'I' when the interface has only one implementation (e.g., `IMenuParser`)
- Use verb-noun combinations for function names (e.g., `ParseMenu`, `GenerateSeedFile`)

## Error Handling
- Always return errors as the last return value
- Use `fmt.Errorf` to add context to errors: `fmt.Errorf("failed to parse menu: %w", err)`
- Create custom error types for specific error cases:
  ```go
  type ValidationError struct {
      Field string
      Error string
  }
  ```
- Handle all errors explicitly; avoid using `_` to ignore errors
- Use error wrapping with `%w` verb for maintaining error chains
- Log errors at the appropriate level (debug/info/error) before returning them

## Data Structuring
### Menu Items
```go
// Preferred structure for menu items
type MenuItem struct {
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Category    string  `json:"category"`
    Options     []Option `json:"options,omitempty"`
}

// Use meaningful field tags
type Option struct {
    Name     string   `json:"name"`
    Choices  []string `json:"choices,omitempty"`
    Price    float64  `json:"price,omitempty"`
}
```

### Seed File Structure
- Place seed data in a package named after the menu type (e.g., `package lunch`, `package dinner`)
- Use descriptive variable names for seed data
- Group related items together
- Include comments describing special cases or requirements

Example:
```go
package dinner

var DinnerMenu = &models.Menu{
    Name: "Evening Dining",
    Items: []models.MenuItem{
        {
            Name:        "Example Item",
            Description: "Example item description...",
            Price:       24.99,
            Category:    "Entrees",
        },
    },
}
```

## Formatting
- Use `gofmt` or `goimports` for consistent formatting
- Maximum line length: 100 characters
- Group imports in the following order:
  1. Standard library
  2. Third-party packages
  3. Internal packages
- Use blank lines to separate logical blocks of code
- Align struct field tags for better readability

## Documentation
- Every exported type, function, or variable must have a comment
- Comments should explain WHY, not WHAT (the code shows what)
- Use complete sentences with proper punctuation
- Include examples for non-obvious functionality

Example:
```go
// ParsePDFMenu extracts structured menu data from a PDF file.
// It handles multi-page menus and supports both single and multi-column layouts.
// Returns an error if the PDF is malformed or cannot be parsed.
func ParsePDFMenu(path string) (*Menu, error)
```

## Testing
- Test files should be named `xxx_test.go`
- Use table-driven tests for multiple test cases
- Test both success and error cases
- Use meaningful test names: `TestParseMenu_ValidPDF`, `TestParseMenu_MalformedInput`

## Version Control
- Use descriptive commit messages
- One logical change per commit
- Reference issue numbers in commits when applicable

## Code Generation
When generating .go seed files:
- Include a comment header indicating the file was generated
- Include the source PDF file name in comments
- Include the generation timestamp
- Use consistent indentation (tabs)
- Sort menu items alphabetically within categories

Example generated file header:
```go
// Code generated by menu-creator. DO NOT EDIT.
// Source: lunch-menu-2024.pdf
// Generated: 2024-01-08 15:30:00
```
