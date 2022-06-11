package strcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/core"
)

var currentTimeCmd = &cobra.Command{
	Use:   "currentTime",
	Short: "Print current time",
	Run: func(cmd *cobra.Command, args []string) {
		util := core.NewCoreUtil()
		fmt.Println(util.Str.CurrentTime())
	},
}
