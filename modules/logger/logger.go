package logger

import (
	"fmt"
	"log"
)

func getzarubaRuntimeName() (name string) {
	name = "🌞 zaruba"
	return fmt.Sprintf("%-12v", name)
}

// Info print info
func Info(format string, v ...interface{}) {
	newFormat := fmt.Sprintf("\033[36m[OUT - %s]\033[0m %s", getzarubaRuntimeName(), format)
	log.Printf(newFormat, v...)
}

// Error print error
func Error(format string, v ...interface{}) {
	newFormat := fmt.Sprintf("\033[31m[ERR - %s]\033[0m %s", getzarubaRuntimeName(), format)
	log.Printf(newFormat, v...)
}
