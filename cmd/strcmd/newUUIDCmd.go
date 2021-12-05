package strcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
)

var newUUIDCmd = &cobra.Command{
	Use:   "newUUID",
	Short: "Generate new UUID string",
	Run: func(cmd *cobra.Command, args []string) {
		util := core.NewCoreUtil()
		fmt.Println(util.Str.NewUUID())
	},
}
