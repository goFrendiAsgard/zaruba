package strcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/dsl"
)

var currentTimeCmd = &cobra.Command{
	Use:   "currentTime",
	Short: "Print current time",
	Run: func(cmd *cobra.Command, args []string) {
		util := dsl.NewDSLUtil()
		fmt.Println(util.Str.CurrentTime())
	},
}
