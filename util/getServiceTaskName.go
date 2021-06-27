package util

import (
	"fmt"

	"github.com/state-alchemists/zaruba/str"
)

func GetServiceTaskName(serviceName string) (taskName string) {
	upperServiceName := str.Capitalize(serviceName)
	taskName = fmt.Sprintf("run%s", upperServiceName)
	return taskName
}
