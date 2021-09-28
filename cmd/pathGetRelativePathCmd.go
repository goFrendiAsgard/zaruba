package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var pathGetRelativePathCmd = &cobra.Command{
	Use:   "getRelativePath <basePath> <targetPath>",
	Short: "Get relative path",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		absBasePath, err := filepath.Abs(args[0])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		absTargetPath, err := filepath.Abs(args[1])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		relPath, err := filepath.Rel(absBasePath, absTargetPath)
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(relPath)
	},
}
