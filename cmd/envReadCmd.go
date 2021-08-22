package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var envReadCmd = &cobra.Command{
	Use:   "read <fileName> [prefix]",
	Short: "Read envmap from file",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 1)
		fileName := args[0]
		envMap, err := godotenv.Read(fileName)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		// cascade prefix
		if len(args) > 1 {
			prefix := args[1]
			envMap = envCascadePrefix(envMap, prefix)
		}
		resultB, err := json.Marshal(envMap)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(string(resultB))
	},
}
