package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zaruba",
	Short: "Task runner and CLI utilities",
	Long: `
 _____                _       _
|__  /__ _ _ __ _   _| |__   / \
  / // _  |  __| | | |  _ \ / _ \
 / /| (_| | |  | |_| | |_) / ___ \
/____\__,_|_|   \__,_|_.__/_/   \_\
Task runner framework and CLI utilities`,
}

var advertisementCmd = &cobra.Command{
	Use:   "advertisement",
	Short: "Advertisement utilities",
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "env utilities",
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

var numCmd = &cobra.Command{
	Use:   "num",
	Short: "Number manipulation utilities",
}

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "path manipulation utilities",
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
	rootCmd.AddCommand(advertisementCmd)
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(linesCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(mapCmd)
	rootCmd.AddCommand(numCmd)
	rootCmd.AddCommand(pathCmd)
	rootCmd.AddCommand(pleaseCmd)
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(strCmd)
	rootCmd.AddCommand(taskCmd)
	rootCmd.AddCommand(utilCmd)

	advertisementCmd.AddCommand(advertisementShowCmd)

	envCmd.AddCommand(envGetMapCmd)

	linesCmd.AddCommand(linesFillCmd)
	linesCmd.AddCommand(linesGetIndexCmd)
	linesCmd.AddCommand(linesIndentCmd)
	linesCmd.AddCommand(linesInsertAfterCmd)
	linesCmd.AddCommand(linesInsertBeforeCmd)
	linesCmd.AddCommand(linesReadCmd)
	linesCmd.AddCommand(linesReplaceCmd)
	linesCmd.AddCommand(linesSubmatchCmd)
	linesCmd.AddCommand(linesWriteCmd)

	listCmd.AddCommand(listAppendCmd)
	listCmd.AddCommand(listContainCmd)
	listCmd.AddCommand(listGetCmd)
	listCmd.AddCommand(listJoinCmd)
	listCmd.AddCommand(listLengthCmd)
	listCmd.AddCommand(listMergeCmd)
	listCmd.AddCommand(listSetCmd)
	listCmd.AddCommand(listValidateCmd)

	mapCmd.AddCommand(mapGetCmd)
	mapCmd.AddCommand(mapGetKeysCmd)
	mapCmd.AddCommand(mapMergeCmd)
	mapCmd.AddCommand(mapSetCmd)
	mapCmd.AddCommand(mapValidateCmd)

	numCmd.AddCommand(numValidateIntCmd)

	pathCmd.AddCommand(pathGetEnvCmd)
	pathCmd.AddCommand(pathGetPortConfigCmd)
	pathCmd.AddCommand(pathGetServiceNameCmd)
	pathCmd.AddCommand(pathGetRelativePathCmd)

	projectCmd.AddCommand(projectAddTaskCmd)
	projectCmd.AddCommand(projectIncludeCmd)
	projectCmd.AddCommand(projectSetValueCmd)
	projectCmd.AddCommand(projectShowLogCmd)
	projectCmd.AddCommand(projectSyncEnvCmd)
	projectCmd.AddCommand(projectSyncEnvFilesCmd)

	strCmd.AddCommand(strAddPrefixCmd)
	strCmd.AddCommand(strGetIndentationCmd)
	strCmd.AddCommand(strIndentCmd)
	strCmd.AddCommand(strNewUUIDCmd)
	strCmd.AddCommand(strRepeatCmd)
	strCmd.AddCommand(strReplaceCmd)
	strCmd.AddCommand(strSplitCmd)
	strCmd.AddCommand(strSubmatchCmd)
	strCmd.AddCommand(strToCamelCmd)
	strCmd.AddCommand(strToKebabCmd)
	strCmd.AddCommand(strToLowerCmd)
	strCmd.AddCommand(strToPascalCmd)
	strCmd.AddCommand(strToSnakeCmd)
	strCmd.AddCommand(strToUpperCmd)

	taskCmd.AddCommand(taskAddDependencyCmd)
	taskCmd.AddCommand(taskAddParentCmd)
	taskCmd.AddCommand(taskIsExistCmd)
	taskCmd.AddCommand(taskSetConfigCmd)
	taskCmd.AddCommand(taskSetEnvCmd)
	taskCmd.AddCommand(taskSyncEnvCmd)

	utilCmd.AddCommand(utilGenerateCmd)
	utilCmd.AddCommand(utilServeCmd)
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
