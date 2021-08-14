package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zaruba",
	Short: "Declarative task runner framework and CLI utilities",
	Long: `
 _____                _       _
|__  /__ _ _ __ _   _| |__   / \
  / // _  |  __| | | |  _ \ / _ \
 / /| (_| | |  | |_| | |_) / ___ \
/____\__,_|_|   \__,_|_.__/_/   \_\
Declarative task runner framework and CLI utilities

Zaruba helps you declare/execute tasks to manage your projects.
Try:
  zaruba please`,
}

var linesCmd = &cobra.Command{
	Use:   "lines",
	Short: "Lines manipulation utilities",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List manipulation utilities",
}

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "Map manipulation utilities",
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Project manipulation utilities",
}

var strCmd = &cobra.Command{
	Use:   "str",
	Short: "String manipulation utilities",
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task manipulation utilities",
}

var utilCmd = &cobra.Command{
	Use:   "util",
	Short: "Utilities",
}

func init() {
	rootCmd.AddCommand(linesCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(mapCmd)
	rootCmd.AddCommand(pleaseCmd)
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(strCmd)
	rootCmd.AddCommand(taskCmd)
	rootCmd.AddCommand(utilCmd)

	projectCmd.AddCommand(projectIncludeCmd)

	strCmd.AddCommand(strAddPrefixCmd)
	strCmd.AddCommand(strGetIndentationCmd)
	strCmd.AddCommand(strIndentCmd)
	strCmd.AddCommand(strRepeatCmd)
	strCmd.AddCommand(strReplaceCmd)
	strCmd.AddCommand(strSubmatchCmd)

	taskCmd.AddCommand(taskAddDependencyCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		decoration := output.NewDecoration()
		logger := output.NewConsoleLogger(decoration)
		logger.Println(err)
		logger.DPrintfError("To run a task you need to type %s%szaruba please <task-name>%s\n", decoration.Bold, decoration.Yellow, decoration.Normal)
		os.Exit(1)
	}
}
