package cmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/toc"
)

var tocLong = `
Create/update documentation directory based on a TOC file.

A TOC file is markdown file representing "table of content" for your documentation.
A tag is written as <!--start[TagName] attribute="value" -->content<!--end[TagName]-->

There are several tagname available for a TOC file:
    - toc: This tag contains list of bulleted-items that will be rendered into documentation structure.
    - code: This tag has several attributes:
        - lang: code language (e.g., python, javascript, bash)
        - src: location of your source code file, relative to TOC file directory
        - cmd: Command to run your code.
      Your code tag will be rendered into a markdown.    

There are also additional tagName for your documentation files:
    - tocHeader: This tag will be filled with documentation header
    - tocSubtopic: This tag will be filled with documentation subtopics
`

var tocExample = `
> cat README.md
# My Cool Project

This is a documentation for my cool project.

<!--startToc-->
- Getting Started
- Concepts
    - Model
    - View
    - Controller
<!--endToc-->

> zaruba toc README.md
`

var tocCmd = &cobra.Command{
	Use:     "toc <tocFileLocation>",
	Short:   "Create/update documentation directory based on a TOC file",
	Long:    tocLong,
	Example: tocExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		tocFile := args[0]
		toc, err := toc.NewToc(tocFile)
		if err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
		if err := toc.RenderNewContent(); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
