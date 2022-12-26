package toc

import (
	"testing"

	"github.com/state-alchemists/zaruba/dsl"
)

var expectedTocFileContent = `# A simple documentation

In this documentation, there will be several things available

- Nested structure
- New links
- Renamed links
- Old links

Before Toc
<!--startToc-->
- [Vehicles](vehicles/README.md)
  - [Motorbike](vehicles/motorbike.md)
  - [Car](vehicles/car.md)
- [Food and Drinks](foodAndDrinks/README.md)
  - [Food](foodAndDrinks/food/README.md)
    - [Fruits](foodAndDrinks/food/fruits/README.md)
      - [Apple](foodAndDrinks/food/fruits/apple.md)
      - [Banana](foodAndDrinks/food/fruits/banana.md)
    - [Nasi Goreng](foodAndDrinks/food/nasiGoreng.md)
  - [Drinks](foodAndDrinks/drinks/README.md)
    - [Coffee](foodAndDrinks/drinks/coffee.md)
<!--endToc-->
After Toc`

func TestNewToc(t *testing.T) {
	util := dsl.NewDSLUtil()
	tocFilePath := "../test-resources/toc/playground/README.md"
	toc, err := NewToc(tocFilePath)
	if err != nil {
		t.Error(err)
		return
	}
	if err := toc.RenderNewContent(); err != nil {
		t.Error(err)
		return
	}
	tocFileContent, err := util.File.ReadText(tocFilePath)
	if err != nil {
		t.Error(err)
		return
	}
	if tocFileContent != expectedTocFileContent {
		t.Errorf("Expected: %s\nActual: %s", expectedTocFileContent, tocFileContent)
	}
}
