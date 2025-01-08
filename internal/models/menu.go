package models

// Menu represents a restaurant menu with multiple items
type Menu struct {
	Name  string
	Items []MenuItem
}

// MenuItem represents a single item in a menu
type MenuItem struct {
	Name        string
	Description string
	Price       float64
	Category    string
}
