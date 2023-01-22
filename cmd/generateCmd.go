package cmd

import (
	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/dsl"
	"github.com/state-alchemists/zaruba/output"
)

var generateExample = `
> mkdir -p template

# template/index.html
> echo '<h1>title</h1>' > template/index.html
> echo 'content' >> template/index.html
# template/start.sh
> echo 'python -m http.server port' > template/start.sh

> zaruba generate template web1 '{"port":3000, "title":"MyWeb", "content":"<p>Hello world!</p>"}'

> ls web1
index.html  start.sh
# web1/index.html
> cat web1/index.html
<h1>MyWeb</h1>
<p>Hello world!</p>
# web1/start.html
> cat web1/start.sh
python -m http.server 3000

> zaruba generate template web2 port 8000 title MySecondWeb content '<p>Hello world!</p>'

> ls web2
index.html  start.sh
# web2/index.html
> cat web2/index.html
<h1>MySecondWeb</h1>
<p>Hello world!</p>
# web2/start.html
> cat web2/start.sh
python -m http.server 8000
`

var generateCmd = &cobra.Command{
	Use:     "generate <sourceTemplatePath> <destinationPath> [{<jsonMapReplacement> | <key> <value>}]",
	Short:   "Generate a directory based on sourceTemplate and jsonMapReplacement",
	Example: generateExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 2)
		sourceTemplatePath, destinationPath := args[0], args[1]
		jsonMapReplacement := "{}"
		if len(args) > 2 {
			var err error
			jsonMapReplacement, err = cmdHelper.ArgToJsonReplacementMap(args, 2)
			if err != nil {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
		}
		util := dsl.NewDSLUtil()
		if err := util.File.Generate(sourceTemplatePath, destinationPath, jsonMapReplacement); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
