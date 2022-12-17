package advertisementcmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/state-alchemists/zaruba/advertisement"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var showLong = `
Show advertisement based on time-pattern in an advertisementFile.

An advertisementFile is a YAML file containing some key-value configurations:

    <advertisement_name>:
      pattern: <YYYY-MM-DD>
      message: <message you want to show>

Please refer to ${ZARUBA_HOME}/advertisement.yaml for more details.
`

var showExample = `
> zaruba advertisement show ~/.zaruba/advertisement.yaml
`

var showCmd = &cobra.Command{
	Use:     "show <strAdvertisementFile>",
	Short:   "Show advertisement based on advertisementFile",
	Long:    showLong,
	Example: showExample,
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 1)
		strAdvertisementFile := args[0]
		advs, err := advertisement.NewAdvs(strAdvertisementFile)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(advs.Get())
	},
}
