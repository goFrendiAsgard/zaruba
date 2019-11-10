package example

import (
	"fmt"
)

// Greet someone politely using golang
func Greet(name string) string {
	return fmt.Sprintf("Hi %s, greeting from golang", name)
}
