package iconer

// Generator is icon generator
type Generator struct {
	iconList []string
	index    int
}

// NewGenerator create new icon generator
func NewGenerator() *Generator {
	return &Generator{
		iconList: []string{"🍏", "🍎", "🍌", "🍉", "🍇", "🍐", "🍊", "🍋", "🍓", "🍈", "🍒", "🍑", "🍍", "🥝", "🍅", "🍆", "🥑"},
		index:    0,
	}
}

// Create new icon
func (g *Generator) Create() string {
	if g.index >= len(g.iconList) {
		g.index = 0
	}
	icon := g.iconList[g.index]
	g.index++
	return icon
}
