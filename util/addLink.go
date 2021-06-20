package util

import (
	"fmt"
)

func AddLink(fileName, source, destination string) (err error) {
	if source == "" {
		return fmt.Errorf("source cannot be empty")
	}
	if destination == "" {
		return fmt.Errorf("destination cannot be empty")
	}
	return SetProjectValue(fileName, fmt.Sprintf("link::%s", destination), source)
}
