package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	commonHelper "github.com/state-alchemists/zaruba/helper"
	"github.com/state-alchemists/zaruba/output"
)

func Exit(cmd *cobra.Command, logger output.Logger, decoration *output.Decoration, err error) {
	if err != nil {
		// get nodeCmd and commandName
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
		// get cmdArgs
		cmdArgs := cmd.Flags().Args()
		jsonCmdArgsBytes, _ := json.Marshal(cmdArgs)
		// get cmdFlagMap
		cmdFlagMap := map[string]string{}
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			cmdFlagMap[flag.Name] = flag.Value.String()
		})
		jsonCmdFlagBytes, _ := json.Marshal(cmdFlagMap)
		// get info
		info := strings.Join([]string{
			fmt.Sprintf("%s %s%sCommand   : %s%s", decoration.ErrorIcon, decoration.Bold, decoration.Red, commandName, decoration.Normal),
			fmt.Sprintf("%s %s%sArguments : %s%s", decoration.ErrorIcon, decoration.Bold, decoration.Red, string(jsonCmdArgsBytes), decoration.Normal),
			fmt.Sprintf("%s %s%sFlags     : %s%s", decoration.ErrorIcon, decoration.Bold, decoration.Red, string(jsonCmdFlagBytes), decoration.Normal),
			fmt.Sprintf("%s %s%sStderr    : %s%s", decoration.ErrorIcon, decoration.Bold, decoration.Red, err.Error(), decoration.Normal),
		}, "\n")
		// print and exit
		logger.Fprintf(os.Stderr, "%s\n", info)
		os.Exit(1)
	}
}

func CheckMinArgCount(cmd *cobra.Command, logger output.Logger, decoration *output.Decoration, args []string, minArgCount int) {
	if len(args) < minArgCount {
		usage := cmd.UsageString()
		argsJsonBytes, _ := json.Marshal(args)
		err := fmt.Errorf("expecting %d arguments, get %d: %s\n%s", minArgCount, len(args), string(argsJsonBytes), usage)
		Exit(cmd, logger, decoration, err)
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

func GetProjectRelFilePath(args []string, argIndex int, defaultFileNames ...string) (filePath string, err error) {
	if len(args) > argIndex {
		return filepath.Abs(args[argIndex])
	}
	workingPath, err := commonHelper.GetWorkingProjectPath()
	if err != nil {
		return "", err
	}
	for index, defaultFileName := range defaultFileNames {
		filePath = filepath.Join(workingPath, defaultFileName)
		if _, err := os.Stat(filePath); err == nil || index == len(defaultFileNames)-1 {
			return filePath, nil
		}
	}
	return "", fmt.Errorf("no matching file: %#v", defaultFileNames)
}

func GetCsvRecordLogger(projectDir string) (csvRecordLogger *output.CSVRecordLogger) {
	logFileName := filepath.Join(projectDir, "logs", "log.zaruba.csv")
	backupFileNameTemplate := "log-%s.zaruba.csv"
	maxLogFileSize, err := strconv.Atoi(os.Getenv("ZARUBA_MAX_LOG_FILE_SIZE"))
	if err != nil {
		maxLogFileSize = 5 * 1024 * 1024
	}
	return output.NewCSVRecordLogger(logFileName, backupFileNameTemplate, maxLogFileSize)
}
