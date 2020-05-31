package example

import (
	"fmt"
	"strings"
)

// Greet greet one person, or the world
func Greet(name string) string {
	if name == "" {
		return "Hello world !!!"
	}
	return fmt.Sprintf("Hello %s", name)
}

// GreetEveryone greet many persons, or everyone
func GreetEveryone(names []string) string {
	if len(names) == 0 {
		return "Hello everyone !!!"
	}
	return fmt.Sprintf("Hello %s, and everyone", strings.Join(names, ", "))
}
