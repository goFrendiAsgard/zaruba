package util

import "strings"

func GetTaskServiceName(taskName string) (serviceName string) {
	if strings.HasPrefix(taskName, "run") && taskName != "run" {
		upperServiceName := strings.TrimPrefix(taskName, "run")
		return strings.ToLower(string(upperServiceName[0])) + upperServiceName[1:]
	}
	return taskName
}
