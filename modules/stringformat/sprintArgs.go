package stringformat

import (
	"fmt"
	"strings"
)

// SprintArgs print arguments for logging
func SprintArgs(args []string) string {
	if len(args) > 0 {
		return fmt.Sprintf("with %d argument(s): %s", len(args), strings.Join(args, ", "))
	}
	return "without argument"
}
