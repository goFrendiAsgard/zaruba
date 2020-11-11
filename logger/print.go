package logger

import "fmt"

var d = NewDecoration()

func printf(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("%s💀%s %s", d.Faint, d.Normal, template)
	return fmt.Printf(pTemplate, args...)
}

// Printf is a wrapper for fmt.Printf
func Printf(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("   %s", template)
	return printf(pTemplate, args...)
}

// PrintfSuccess is a wrapper for fmt.Printf
func PrintfSuccess(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("🎉 %s", template)
	return printf(pTemplate, args...)
}

// PrintfError is a wrapper for fmt.Printf
func PrintfError(template string, args ...interface{}) (n int, err error) {

	pTemplate := fmt.Sprintf("🔥 %s", template)
	return printf(pTemplate, args...)
}

// PrintfStarted is a wrapper for fmt.Printf
func PrintfStarted(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("🏁 %s", template)
	return printf(pTemplate, args...)
}

// PrintfKill is a wrapper for fmt.Printf
func PrintfKill(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("🔪 %s", template)
	return printf(pTemplate, args...)
}

// PrintfInspect is a wrapper for fmt.Printf
func PrintfInspect(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("🔎 %s", template)
	return printf(pTemplate, args...)
}
