package projectcmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	cmdHelper "github.com/state-alchemists/zaruba/cmd/helper"
	"github.com/state-alchemists/zaruba/output"
)

var showLogCmd = &cobra.Command{
	Use:   "showLog <taskNamePattern> [logFile]",
	Short: "Show log",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		cmdHelper.CheckMinArgCount(cmd, logger, decoration, args, 0)
		taskNamePattern := ".*"
		if len(args) > 0 {
			taskNamePattern = args[0]
		}
		logFileName := "logs/log.zaruba.csv"
		if len(args) > 1 {
			logFileName = args[1]
		}
		message, err := getLog(decoration, taskNamePattern, logFileName)
		if err != nil {
			cmdHelper.Exit(cmd, args, logger, decoration, err)
		}
		fmt.Println(message)
	},
}

func getLog(decoration *output.Decoration, pattern, logFileName string) (logMessage string, err error) {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return "", err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return "", err
	}
	logMessage = ""
	for _, row := range records {
		timestamp, output_type, command_type, taskName, log, sessionId := row[0], row[1], row[2], row[3], row[4], row[5]
		sessionIdParts := strings.Split(sessionId, "-")
		shortSessionId := fmt.Sprintf("%s-%s", sessionIdParts[0], sessionIdParts[1])
		for len(shortSessionId) < 30 {
			shortSessionId += " "
		}
		match, err := regexp.Match(pattern, []byte(taskName))
		if err != nil || !match {
			continue
		}
		// adjust commandType
		if command_type == "START" {
			command_type = decoration.RunIcon
		} else {
			command_type = decoration.InspectIcon
		}
		// adjust output_type
		if output_type == "ERR" {
			output_type = decoration.ErrorIcon
		} else {
			output_type = decoration.EmptyIcon
		}
		logMessage += fmt.Sprintf("%s%s%s \t %s %s %s\t%s%s%s %s\n", decoration.Faint, shortSessionId, decoration.Normal, output_type, command_type, taskName, decoration.Faint, timestamp[:23], decoration.Normal, log)
	}
	return logMessage, nil
}
