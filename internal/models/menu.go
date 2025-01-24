package models

import (
	"fmt"
	"regexp"
	"strings"
)

// Location represents a restaurant or venue
type Location struct {
	ID       string
	Name     string
	Address  string
	Menus    []Menu
}

// Menu represents a collection of menu categories
type Menu struct {
	ID         string
	Name       string
	LocationID string
	Categories []Category
}

// Category represents a group of related menu items
type Category struct {
	ID          string
	Name        string
	Description string
	MenuID      string
	Items       []MenuItem
}

// MenuItem represents a single item in a menu category
type MenuItem struct {
	ID          string
	Name        string
	Description string
	Price       float64
	CategoryID  string
	Options     []Option
}

// Option represents a group of choices for a menu item
type Option struct {
	ID          string
	Name        string
	Description string
	MinSelect   int
	MaxSelect   int
	ItemID      string
	Items       []OptionItem
}

// OptionItem represents a single choice within an option group
type OptionItem struct {
	Name        string
	Description string
	Price       float64
}

// Helper functions for generating IDs and normalizing names
func GenerateID(name string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	parts := reg.Split(name, -1)
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.Title(strings.ToLower(parts[i]))
		}
	}
	if len(parts) > 0 {
		parts[0] = strings.ToLower(parts[0])
	}
	return strings.Join(parts, "")
}

// NormalizeCategory standardizes category names
func NormalizeCategory(category string) string {
	// Default category mappings - can be extended via configuration
	categoryMap := map[string]string{
		"STARTERS":     "Appetizers",
		"APPETIZERS":   "Appetizers",
		"SMALL PLATES": "Appetizers",
		"ENTREES":      "Main Courses",
		"MAINS":        "Main Courses",
		"DESSERTS":     "Desserts",
		"BEVERAGES":    "Beverages",
	}

	if normalized, ok := categoryMap[strings.ToUpper(category)]; ok {
		return normalized
	}
	return strings.Title(strings.ToLower(category))
}

// CreateLocation creates a new Location with the given name and address
func CreateLocation(name, address string) *Location {
	return &Location{
		ID:      GenerateID(name),
		Name:    name,
		Address: address,
		Menus:   []Menu{},
	}
}

// CreateMenu creates a new Menu with the given name and location ID
func CreateMenu(name, displayName, locationID string) *Menu {
	return &Menu{
		ID:         GenerateID(fmt.Sprintf("%s%s", name, displayName)),
		Name:       displayName,
		LocationID: locationID,
		Categories: []Category{},
	}
}

// CreateCategory creates a new Category with the given name and menu ID
func CreateCategory(name, description, menuID string) *Category {
	return &Category{
		ID:          GenerateID(name),
		Name:        name,
		Description: description,
		MenuID:      menuID,
		Items:       []MenuItem{},
	}
}

// CreateItem creates a new MenuItem with the given details
func CreateItem(name, description string, price float64, categoryID string) *MenuItem {
	return &MenuItem{
		ID:          GenerateID(name),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		Options:     []Option{},
	}
}

// CreateOption creates a new Option with the given details and option items
func CreateOption(name, description string, minSelect, maxSelect int, itemID string, items []OptionItem) *Option {
	return &Option{
		ID:          GenerateID(name),
		Name:        name,
		Description: description,
		MinSelect:   minSelect,
		MaxSelect:   maxSelect,
		ItemID:      itemID,
		Items:       items,
	}
}
