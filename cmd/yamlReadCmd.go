package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
	"gopkg.in/yaml.v3"
)

var yamlReadCmd = &cobra.Command{
	Use:   "read <fileName>",
	Short: "Read yaml from file",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		fileName := args[0]
		util := core.NewUtil()
		yamlScript, err := util.File.ReadText(fileName)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		// get yaml
		var interfaceContent interface{}
		if err = yaml.Unmarshal([]byte(yamlScript), &interfaceContent); err != nil {
			exit(cmd, logger, decoration, err)
		}
		interfaceContent = convertYamlObj(interfaceContent)
		resultB, err := json.Marshal(interfaceContent)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
