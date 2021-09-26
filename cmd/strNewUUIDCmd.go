package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/utility"
)

var strNewUUIDCmd = &cobra.Command{
	Use:   "newUUID",
	Short: "Generate new UUID string",
	Run: func(cmd *cobra.Command, args []string) {
		util := utility.NewUtil()
		fmt.Println(util.Str.NewUUID())
	},
}
