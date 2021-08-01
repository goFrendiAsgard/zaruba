package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var getRelativePathCmd = &cobra.Command{
	Use:   "getRelativePath <basePath> <targetPath>",
	Short: "Get value or default value",
	Run: func(cmd *cobra.Command, args []string) {
		commandName := cmd.Name()
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(commandName, logger, decoration, args, 2)
		absBasePath, err := filepath.Abs(args[0])
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		absTargetPath, err := filepath.Abs(args[1])
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		relPath, err := filepath.Rel(absBasePath, absTargetPath)
		if err != nil {
			exit(commandName, logger, decoration, err)
		}
		fmt.Println(relPath)
	},
}

func init() {
	rootCmd.AddCommand(getRelativePathCmd)
}
