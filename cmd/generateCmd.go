package cmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var generateExample = `
> ls template
index.html  start.sh

> cat template/index.html
<h1>title</h1>
content

> cat template/start.sh
python -m http.server port

> zaruba generate template web '{"port":3000, "title":"MyWeb", "content":"<p>Hello world!</p>"}'

> ls web
index.html  start.sh

> cat web/index.html
<h1>MyWeb</h1>
<p>Hello world!</p>

> cat web/start.sh
python -m http.server 3000
`

var generateCmd = &cobra.Command{
	Use:     "generate <sourceTemplatePath> <destinationPath> [jsonMapReplacement]",
	Short:   "Generate a directory based on sourceTemplate and jsonMapReplacement",
	Example: generateExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		sourceTemplatePath, destinationPath := args[0], args[1]
		jsonMapReplacement := "{}"
		if len(args) > 2 {
			jsonMapReplacement = args[2]
		}
		util := dsl.NewDSLUtil()
		if err := util.File.Generate(sourceTemplatePath, destinationPath, jsonMapReplacement); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
