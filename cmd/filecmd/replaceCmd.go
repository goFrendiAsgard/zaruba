package filecmd

import (
	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var replaceExample = `
> echo 'Capital of country is city' > example1.txt
> zaruba str replace example1.txt '{"country": "Indonesia", "city": "Jakarta"}'
> cat example1.txt
Capital of Indonesia is Jakarta

> echo 'Capital of country is city' > example2.txt
> zaruba str replace example2.txt country Japan city Tokyo
> cat example2.txt
Capital of Japan is Tokyo

> echo "def add(a):" > example.py
> echo "    pass" >> example.py
> echo "" >> example.py
> echo "def minus(" >> example.py
> echo "    a" >> example.py
> echo "):" >> example.py
> echo "    pass" >> example.py
> echo "" >> example.py
> echo "class Something():" >> example.py
> echo "    def __init__(a):" >> example.py
> echo "        pass" >> example.py
> zaruba file replace example.py '(?U)(?m)(?s)def (.*)\((.*)([\n\t ]*)\):' 'def $1($2, b$3):'
> cat example.py
def add(a, b):
    pass

def minus(
    a, b
):
    pass

class Something():
    def __init__(a, b):
        pass
`

var replaceCmd = &cobra.Command{
	Use:     "replace <strFileName> [{<jsonMapReplacement> | <key> <value>}]",
	Short:   "Replace string by jsonMapReplacement",
	Example: replaceExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		strFileName := args[0]
		jsonMapReplacement := "{}"
		if len(args) > 1 {
			var err error
			jsonMapReplacement, err = cmdHelper.ArgToJsonReplacementMap(args, 1)
			if err != nil {
				cmdHelper.Exit(cmd, logger, decoration, err)
			}
		}
		util := dsl.NewDSLUtil()
		if err := util.File.Replace(strFileName, jsonMapReplacement); err != nil {
			cmdHelper.Exit(cmd, logger, decoration, err)
		}
	},
}
