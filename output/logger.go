package output

import (
	"fmt"
	"io"
)

type Logger interface {
	Print(args ...interface{}) (n int, err error)
	Println(args ...interface{}) (n int, err error)
	Printf(template string, args ...interface{}) (n int, err error)
	Fprint(w io.Writer, a ...interface{}) (n int, err error)
	Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	Fprintf(w io.Writer, template string, a ...interface{}) (n int, err error)
	DPrintf(template string, args ...interface{}) (n int, err error)
	DPrintfSuccess(template string, args ...interface{}) (n int, err error)
	DPrintfError(template string, args ...interface{}) (n int, err error)
	DPrintfStarted(template string, args ...interface{}) (n int, err error)
	DPrintfKill(template string, args ...interface{}) (n int, err error)
	DPrintfInspect(template string, args ...interface{}) (n int, err error)
}

type ConsoleLogger struct {
	d *Decoration
}

// NewConsoleLogger create new output
func NewConsoleLogger(decoration *Decoration) *ConsoleLogger {
	return &ConsoleLogger{
		d: decoration,
	}
}

// Print basically fmt.Print
func (l *ConsoleLogger) Print(args ...interface{}) (n int, err error) {
	return fmt.Print(args...)
}

// Println basically fmt.Println
func (l *ConsoleLogger) Println(args ...interface{}) (n int, err error) {
	return fmt.Println(args...)
}

// Printf bascically fmt.Printf
func (l *ConsoleLogger) Printf(template string, args ...interface{}) (n int, err error) {
	return fmt.Printf(template, args...)
}

func (l *ConsoleLogger) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprint(w, a...)
}

func (l *ConsoleLogger) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, a...)
}

func (l *ConsoleLogger) Fprintf(w io.Writer, template string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, template, a...)
}

func (l *ConsoleLogger) dPrintf(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("%s %s", l.d.Skull, template)
	return fmt.Printf(pTemplate, args...)
}

// DPrintf is a decorated fmt.Printf
func (l *ConsoleLogger) DPrintf(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("%s %s", l.d.Empty, template)
	return l.dPrintf(pTemplate, args...)
}

// DPrintfSuccess is a decorated fmt.Printf, indicate success
func (l *ConsoleLogger) DPrintfSuccess(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("%s %s", l.d.Success, template)
	return l.dPrintf(pTemplate, args...)
}

// DPrintfError is a decorated fmt.Printf, indicate error
func (l *ConsoleLogger) DPrintfError(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("%s %s", l.d.Error, template)
	return l.dPrintf(pTemplate, args...)
}

// DPrintfStarted is a decorated fmt.Printf, indicate process started
func (l *ConsoleLogger) DPrintfStarted(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("%s %s", l.d.Start, template)
	return l.dPrintf(pTemplate, args...)
}

// DPrintfKill is a decorated fmt.Printf, indicate process killed
func (l *ConsoleLogger) DPrintfKill(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("%s %s", l.d.Kill, template)
	return l.dPrintf(pTemplate, args...)
}

// DPrintfInspect is a decorated fmt.Printf, indicate process inspection
func (l *ConsoleLogger) DPrintfInspect(template string, args ...interface{}) (n int, err error) {
	pTemplate := fmt.Sprintf("%s %s", l.d.Inspect, template)
	return l.dPrintf(pTemplate, args...)
}
