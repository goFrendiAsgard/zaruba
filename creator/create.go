package creator

import (
	"os"
)

// Create something from template
func Create(template string, target string, interactively bool) error {
	os.MkdirAll(target, os.ModePerm)
	return nil
}
