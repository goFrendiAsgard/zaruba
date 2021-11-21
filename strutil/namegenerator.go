package strutil

import (
	"fmt"
	"math/rand"
	"strings"
)

type NameGenerator struct {
	adjectives []string
	nouns      []string
}

func NewNameGenerator() *NameGenerator {
	return &NameGenerator{
		adjectives: []string{"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark", "summer",
			"icy", "delicate", "quiet", "white", "cool", "spring", "winter", "patient",
			"twilight", "dawn", "crimson", "wispy", "weathered", "blue", "billowing",
			"broken", "cold", "damp", "falling", "frosty", "green", "long", "late", "lingering",
			"bold", "little", "morning", "muddy", "old", "red", "rough", "still", "small",
			"sparkling", "throbbing", "shy", "wandering", "withered", "wild", "black",
			"young", "holy", "solitary", "fragrant", "aged", "snowy", "proud", "floral",
			"restless", "divine", "polished", "ancient", "purple", "lively", "nameless"},
		nouns: []string{"waterfall", "river", "breeze", "moon", "rain", "wind", "sea", "morning",
			"snow", "lake", "sunset", "pine", "shadow", "leaf", "dawn", "glitter", "forest",
			"hill", "cloud", "meadow", "sun", "glade", "bird", "brook", "butterfly",
			"bush", "dew", "dust", "field", "fire", "flower", "firefly", "feather", "grass",
			"haze", "mountain", "night", "pond", "darkness", "snowflake", "silence",
			"sound", "sky", "shape", "surf", "thunder", "violet", "water", "wildflower",
			"wave", "water", "resonance", "sun", "wood", "dream", "cherry", "tree", "fog",
			"frost", "voice", "paper", "frog", "smoke", "star"},
	}
}

func (generator *NameGenerator) Generate() string {
	r := rand.New(rand.New(rand.NewSource(99)))
	adjective := generator.adjectives[r.Intn(len(generator.adjectives))]
	noun := strings.Title(generator.adjectives[r.Intn(len(generator.nouns))])
	return fmt.Sprintf("%v%v", adjective, noun)
}
