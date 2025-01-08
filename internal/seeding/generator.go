package seeding

import (
	"fmt"
	"os"
)

// GenerateSeedFile demonstrates how we might create a .go file for seeding
func GenerateSeedFile(menuName string) error {
	file, err := os.Create(fmt.Sprintf("%s_seed.go", menuName))
	if err != nil {
		return fmt.Errorf("failed to create seed file: %w", err)
	}
	defer file.Close()
	// TODO: Write actual .go code contents here
	return nil
}
