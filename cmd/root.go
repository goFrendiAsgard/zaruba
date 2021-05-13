package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zaruba",
	Short: "Declarative Task Runner Framework",
	Long: `💀 Declarative Task Runner Framework

Zaruba help you execute tasks faster and easier.
Try:
  zaruba please`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		logger.Println(err)
		logger.DPrintfError("Do you mean %s%szaruba please%s?\n", decoration.Bold, decoration.Yellow, decoration.Normal)
		os.Exit(1)
	}
}
