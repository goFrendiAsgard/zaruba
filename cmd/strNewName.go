package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
)

var strNewName = &cobra.Command{
	Use:   "newName",
	Short: "Generate new name",
	Run: func(cmd *cobra.Command, args []string) {
		util := core.NewCoreUtil()
		fmt.Println(util.Str.NewName())
	},
}
