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
- [Food and Drinks](food-and-drinks/README.md)
  - [Food](food-and-drinks/food/README.md)
    - [Fruits](food-and-drinks/food/fruits/README.md)
      - [Apple](food-and-drinks/food/fruits/apple.md)
      - [Banana](food-and-drinks/food/fruits/banana.md)
    - [Nasi Goreng](food-and-drinks/food/nasi-goreng.md)
  - [Drinks](food-and-drinks/drinks/README.md)
    - [Coffee](food-and-drinks/drinks/coffee.md)
    - [Tea](food-and-drinks/drinks/tea.md)
<!--endToc-->
After Toc`

var expectedFoodFileContent = `<!--startTocHeader-->
[🏠](../../README.md) > [Food and Drinks](../README.md)
# Food
<!--endTocHeader-->

Article about food

<!--startTocSubtopic-->
- [Fruits](fruits/README.md)
  - [Apple](fruits/apple.md)
  - [Banana](fruits/banana.md)
- [Nasi Goreng](nasi-goreng.md)
<!--endTocSubtopic-->`

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
	// test TOC file content
	tocFileContent, err := util.File.ReadText(tocFilePath)
	if err != nil {
		t.Error(err)
		return
	}
	if tocFileContent != expectedTocFileContent {
		t.Errorf("Expected: %s\nActual: %s", expectedTocFileContent, tocFileContent)
	}
	// food content
	foodFileContent, err := util.File.ReadText("../test-resources/toc/playground/food-and-drinks/food/README.md")
	if err != nil {
		t.Error(err)
		return
	}
	if foodFileContent != expectedFoodFileContent {
		t.Errorf("Expected: %s\nActual: %s", expectedFoodFileContent, foodFileContent)
	}
}
