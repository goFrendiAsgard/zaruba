package iconer

// Generator is icon generator
type Generator struct {
	IconList []string
	Index    int
}

// NewGenerator create new icon generator
func NewGenerator() *Generator {
	return &Generator{
		IconList: []string{"🍏", "🍎", "🍌", "🍉", "🍇", "🍐", "🍊", "🍋", "🍓", "🍈", "🍒", "🍑", "🍍", "🥝", "🍅", "🍆", "🥑"},
		Index:    0,
	}
}

// Create new icon
func (g *Generator) Create() string {
	if g.Index >= len(g.IconList) {
		g.Index = 0
	}
	icon := g.IconList[g.Index]
	g.Index++
	return icon
}
