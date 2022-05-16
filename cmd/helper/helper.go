package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

func Exit(cmd *cobra.Command, args []string, logger output.Logger, decoration *output.Decoration, err error) {
	if err != nil {
		nodeCmd := cmd
		commandName := ""
		for nodeCmd != nil {
			if commandName == "" {
				commandName = nodeCmd.Name()
			} else {
				commandName = fmt.Sprintf("%s %s", nodeCmd.Name(), commandName)
			}
			nodeCmd = nodeCmd.Parent()
		}
		argsJsonBytes, _ := json.Marshal(args)
		commandInfo := fmt.Sprintf("%s %s%sCommand   : %s%s", decoration.ErrorIcon, decoration.Bold, decoration.Red, commandName, decoration.Normal)
		argumentInfo := fmt.Sprintf("%s %s%sArguments : %s%s", decoration.ErrorIcon, decoration.Bold, decoration.Red, string(argsJsonBytes), decoration.Normal)
		errorInfo := fmt.Sprintf("%s %s%sStderr    : %s%s", decoration.ErrorIcon, decoration.Bold, decoration.Red, err.Error(), decoration.Normal)
		logger.Fprintf(os.Stderr, "%s\n%s\n%s\n", commandInfo, argumentInfo, errorInfo)
		os.Exit(1)
	}
}

func CheckMinArgCount(cmd *cobra.Command, logger output.Logger, decoration *output.Decoration, args []string, minArgCount int) {
	if len(args) < minArgCount {
		usage := cmd.UsageString()
		argsJsonBytes, _ := json.Marshal(args)
		err := fmt.Errorf("expecting %d arguments, get %d: %s\n%s", minArgCount, len(args), string(argsJsonBytes), usage)
		Exit(cmd, args, logger, decoration, err)
	}
}

func GetDecoration(decorationMode string) (decoration *output.Decoration) {
	switch decorationMode {
	case "plain":
		return output.NewPlainDecoration()
	case "colorless":
		return output.NewColorlessDecoration()
	default:
		return output.NewDefaultDecoration()
	}
}

func GetCsvRecordLogger(projectDir string) (csvRecordLogger *output.CSVRecordLogger) {
	logFile := filepath.Join(projectDir, "log.zaruba.csv")
	return output.NewCSVRecordLogger(logFile)
}
