package strcmd

import (
	"fmt"

	"github.com/state-alchemists/zaruba/dsl"

	"github.com/spf13/cobra"
)

var newNameCmd = &cobra.Command{
	Use:   "newName",
	Short: "Generate new name",
	Run: func(cmd *cobra.Command, args []string) {
		util := dsl.NewDSLUtil()
		fmt.Println(util.Str.NewName())
	},
}
