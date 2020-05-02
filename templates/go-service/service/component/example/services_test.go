package example

import "testing"

// TestGreetEmptyParameter test greet function
func TestGreetEmptyParameter(t *testing.T) {
	greetings := Greet("")
	if greetings != "Hello world !!!" {
		t.Errorf("Unexpected result: %s", greetings)
	}
}

// TestGreetNonEmptyParameter test greet function
func TestGreetNonEmptyParameter(t *testing.T) {
	greetings := Greet("Kouga")
	if greetings != "Hello Kouga" {
		t.Errorf("Unexpected result: %s", greetings)
	}
}

// TestGreetEveryoneEmptyParameter test greetEveryone function
func TestGreetEveryoneEmptyParameter(t *testing.T) {
	greetings := GreetEveryone([]string{})
	if greetings != "Hello everyone !!!" {
		t.Errorf("Unexpected result: %s", greetings)
	}
}

// TestGreetEveryoneNoneEmptyParameter test greetEveryone function
func TestGreetEveryoneNoneEmptyParameter(t *testing.T) {
	greetings := GreetEveryone([]string{"Kouga", "Kaoru"})
	if greetings != "Hello Kouga, Kaoru, and everyone" {
		t.Errorf("Unexpected result: %s", greetings)
	}
}
