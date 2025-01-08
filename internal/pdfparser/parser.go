package pdfparser

import (
	"fmt"
	"github.com/grow-assistant/menu-creator/internal/models"
)

// ParsePDFMenu extracts menu data from a PDF file's path
func ParsePDFMenu(pdfPath string) (*models.Menu, error) {
	// TODO: use a robust library (e.g., "rsc.io/pdf" or "github.com/pdfcpu/pdfcpu")
	// Then parse to build a Menu struct
	return nil, fmt.Errorf("not implemented")
}
