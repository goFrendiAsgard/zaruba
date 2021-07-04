package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"

	"github.com/state-alchemists/zaruba/output"
)

func GetLogAsString(decoration *output.Decoration, logFile, pattern string) (logMessage string, err error) {
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
