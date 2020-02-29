package file

import (
	"path/filepath"
)

// GetAbsoluteLocation get absolute location
func GetAbsoluteLocation(dirPath, location string) string {
	if filepath.IsAbs(location) {
		return location
	}
	absDirPath, err := filepath.Abs(dirPath)
	if err != nil {
		absDirPath = dirPath
	}
	return filepath.Join(absDirPath, location)
}
