package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/logger"
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
		fmt.Println(err)
		d := logger.NewDecoration()
		logger.PrintfError("Do you mean %s%szaruba please%s?\n", d.Bold, d.Yellow, d.Normal)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
