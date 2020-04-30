package logger

import (
	"fmt"
	"log"
)

func getzarubaRuntimeName() (name string) {
	return "💀 ZARUBA  "
}

// Info print info
func Info(format string, v ...interface{}) {
	newFormat := fmt.Sprintf("\033[36mOUT - %s %s\033[0m", getzarubaRuntimeName(), format)
	log.Printf(newFormat, v...)
}

// Error print error
func Error(format string, v ...interface{}) {
	newFormat := fmt.Sprintf("\033[31mERR - %s %s\033[0m", getzarubaRuntimeName(), format)
	log.Printf(newFormat, v...)
}

// Fatal print error and exit
func Fatal(v ...interface{}) {
	prefix := fmt.Sprintf("\033[31mERR - %s ", getzarubaRuntimeName())
	suffix := "\033[0m"
	newV := []interface{}{prefix}
	newV = append(newV, v...)
	newV = append(newV, suffix)
	log.Fatal(newV...)
}
