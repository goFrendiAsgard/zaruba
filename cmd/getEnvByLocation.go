package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
	"github.com/state-alchemists/zaruba/util"
)

var getEnvByLocationFormat string
var getEnvByLocationCmd = &cobra.Command{
	Use:   "getEnvByLocation <serviceLocation>",
	Short: "Get environment (in JSON format) by location",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		if len(args) < 1 {
			showErrorAndExit(logger, decoration, fmt.Errorf("too few argument for getEnvByLocation"))
		}
		envMap, err := util.GetEnvByLocation(args[0])
		if err != nil {
			showErrorAndExit(logger, decoration, err)
		}
		// output format: json
		if strings.ToLower(getEnvByLocationFormat) == "json" {
			envMapJsonB, err := json.Marshal(envMap)
			if err != nil {
				showErrorAndExit(logger, decoration, err)
			}
			fmt.Println(string(envMapJsonB))
			return
		}
		// output format: not specified
		for key, val := range envMap {
			fmt.Printf("%s=%s\n", key, val)
		}
	},
}

func init() {
	rootCmd.AddCommand(getEnvByLocationCmd)
	getEnvByLocationCmd.Flags().StringVarP(&getEnvByLocationFormat, "format", "f", "", "output format")

}
