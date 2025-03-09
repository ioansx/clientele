package validations

import (
	"fmt"
	"slices"
)

var validPages = []string{
	"bash",
	"cd",
	"grep",
	"ls",
	"man",
	"tar",
}

func ValidateManGet(page string) error {
	if page == "" {
		return fmt.Errorf("'page' is empty")
	}
	if !slices.Contains(validPages[:], page) {
		return fmt.Errorf("'page' must be one of %v", validPages)
	}
	return nil
}
