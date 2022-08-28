package strutil

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type NameGenerator struct {
	adjectives []string
	nouns      []string
}

func NewNameGenerator() *NameGenerator {
	return &NameGenerator{
		adjectives: []string{"autumn", "hidden", "bitter", "misty", "silent",
			"empty", "dry", "dark", "summer", "icy", "delicate", "quiet", "white", "cool",
			"spring", "winter", "patient", "twilight", "dawn", "crimson", "wispy",
			"weathered", "blue", "billowing", "broken", "cold", "damp", "falling",
			"frosty", "green", "long", "late", "lingering", "bold", "little", "morning",
			"muddy", "old", "red", "rough", "still", "small", "sparkling", "throbbing",
			"shy", "wandering", "withered", "wild", "black", "holy", "solitary",
			"fragrant", "aged", "snowy", "proud", "floral", "restless", "divine",
			"polished", "purple", "lively", "nameless", "puffy", "fluffy",
			"calm", "young", "golden", "avenging", "ancestral", "ancient", "argent",
			"reckless", "daunting", "short", "rising", "strong", "timber", "tumbling",
			"silver", "dusty", "celestial", "cosmic", "crescent", "double", "far", "half",
			"inner", "milky", "northern", "southern", "eastern", "western", "outer",
			"terrestrial", "huge", "deep", "epic", "titanic", "mighty", "powerful"},
		nouns: []string{"waterfall", "river", "breeze", "moon", "rain",
			"wind", "sea", "morning", "snow", "lake", "sunset", "pine", "shadow", "leaf",
			"dawn", "glitter", "forest", "hill", "cloud", "meadow", "glade",
			"bird", "brook", "butterfly", "bush", "dew", "dust", "field",
			"flower", "firefly", "feather", "grass", "haze", "mountain", "night", "pond",
			"darkness", "snowflake", "silence", "sound", "sky", "shape", "surf",
			"thunder", "violet", "wildflower", "wave", "water", "resonance",
			"sun", "wood", "dream", "cherry", "tree", "fog", "frost", "voice", "paper",
			"frog", "smoke", "star", "sierra", "castle", "fortress", "tiger", "day",
			"sequoia", "cedar", "wrath", "blessing", "spirit", "nova", "storm", "burst",
			"protector", "drake", "dragon", "knight", "fire", "king", "jungle", "queen",
			"giant", "elemental", "throne", "game", "weed", "stone", "apogee", "bang",
			"cluster", "corona", "cosmos", "equinox", "horizon", "light", "nebula",
			"solstice", "spectrum", "universe", "magnitude", "parallax"},
	}
}

func (generator *NameGenerator) Generate() string {
	r := rand.New(rand.New(rand.NewSource(time.Now().UnixNano())))
	adjective := generator.adjectives[r.Intn(len(generator.adjectives))]
	noun := strings.Title(generator.nouns[r.Intn(len(generator.nouns))])
	return fmt.Sprintf("%v%v", adjective, noun)
}
