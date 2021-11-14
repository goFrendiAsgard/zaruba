package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
	"github.com/state-alchemists/zaruba/output"
)

var pathGetRelativePathCmd = &cobra.Command{
	Use:   "getRelativePath <basePath> <targetPath>",
	Short: "Get relative path",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		basePath := args[0]
		targetPath := args[1]
		util := core.NewCoreUtil()
		relPath, err := util.Path.GetRelativePath(basePath, targetPath)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(relPath)
	},
}
