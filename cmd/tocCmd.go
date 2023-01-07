package cmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/toc"
)

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
	Short:   "Create/update documentations based on TOC file",
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
