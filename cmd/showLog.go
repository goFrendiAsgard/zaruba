package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/state-alchemists/zaruba/output"
)

var projectShowLogCmd = &cobra.Command{
	Use:   "showLog <logFile> <taskNamePattern>",
	Short: "Show log",
	Run: func(cmd *cobra.Command, args []string) {
		decoration := output.NewDefaultDecoration()
		logger := output.NewConsoleLogger(decoration)
		checkMinArgCount(cmd, logger, decoration, args, 2)
		message, err := getLog(decoration, args[0], args[1])
		if err != nil {
			exit(cmd, logger, decoration, err)
		}
		fmt.Println(message)
	},
}

func getLog(decoration *output.Decoration, logFile, pattern string) (logMessage string, err error) {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
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
		timestamp, output_type, command_type, taskName, log := row[0], row[1], row[2], row[3], row[4]
		match, err := regexp.Match(pattern, []byte(taskName))
		if err != nil || !match {
			continue
		}
		// adjust commandType
		if command_type == "START" {
			command_type = decoration.Run
		} else {
			command_type = decoration.Inspect
		}
		// adjust output_type
		if output_type == "ERR" {
			output_type = decoration.Error
		} else {
			output_type = decoration.Empty
		}
		logMessage += fmt.Sprintf("%s %s %s\t%s%s%s %s\n", output_type, command_type, taskName, decoration.Faint, timestamp[:23], decoration.Normal, log)
	}
	return logMessage, nil
}
