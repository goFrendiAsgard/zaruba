package output

import (
	"fmt"
	"io"
	"strings"
	"sync"
)

type MockLoggerData struct {
	Str string
}

func NewMockLoggerData(args ...interface{}) MockLoggerData {
	strArgs := []string{}
	for _, arg := range args {
		strArgs = append(strArgs, fmt.Sprintf("%s", arg))
	}
	return MockLoggerData{Str: strings.Join(strArgs, "")}
}

type MockLoggerTrigger func()

type MockLogger struct {
	Data     []MockLoggerData
	Mutex    *sync.Mutex
	Triggers map[string]MockLoggerTrigger
}

func NewMockLogger() *MockLogger {
	return &MockLogger{
		Data:     []MockLoggerData{},
		Mutex:    &sync.Mutex{},
		Triggers: map[string]MockLoggerTrigger{},
	}
}

func (m *MockLogger) RegisterTrigger(subStr string, trigger MockLoggerTrigger) {
	m.Triggers[subStr] = trigger
}

func (m *MockLogger) GetOutput() (output string) {
	lines := []string{}
	for _, data := range m.Data {
		lines = append(lines, strings.Trim(data.Str, "\n"))
	}
	return strings.Join(lines, "\n")
}

func (m *MockLogger) GetLineIndex(subStr string) (index int) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	for index, data := range m.Data {
		if strings.Contains(data.Str, subStr) {
			return index
		}
	}
	return -1
}

func (m *MockLogger) Print(args ...interface{}) (n int, err error) {
	return m.print(args...)
}

func (m *MockLogger) Println(args ...interface{}) (n int, err error) {
	return m.print(args...)
}

func (m *MockLogger) Printf(template string, args ...interface{}) (n int, err error) {
	return m.printf(template, args...)
}

func (m *MockLogger) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return m.print(a...)
}

func (m *MockLogger) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return m.print(a...)
}

func (m *MockLogger) Fprintf(w io.Writer, template string, a ...interface{}) (n int, err error) {
	return m.printf(template, a...)
}

func (m *MockLogger) DPrintf(template string, args ...interface{}) (n int, err error) {
	return m.printf(template, args...)
}

func (m *MockLogger) DPrintfSuccess(template string, args ...interface{}) (n int, err error) {
	return m.printf(template, args...)
}

func (m *MockLogger) DPrintfError(template string, args ...interface{}) (n int, err error) {
	return m.printf(template, args...)
}

func (m *MockLogger) DPrintfStarted(template string, args ...interface{}) (n int, err error) {
	return m.printf(template, args...)
}

func (m *MockLogger) DPrintfKill(template string, args ...interface{}) (n int, err error) {
	return m.printf(template, args...)
}

func (m *MockLogger) DPrintfInspect(template string, args ...interface{}) (n int, err error) {
	return m.printf(template, args...)
}

func (m *MockLogger) printf(template string, args ...interface{}) (n int, err error) {
	str := fmt.Sprintf(template, args...)
	return m.print(str)
}

func (m *MockLogger) print(args ...interface{}) (n int, err error) {
	m.Mutex.Lock()
	data := NewMockLoggerData(args...)
	m.Data = append(m.Data, data)
	m.Mutex.Unlock()
	for subStr, trigger := range m.Triggers {
		if strings.Contains(data.Str, subStr) {
			trigger()
		}
	}
	return 0, nil
}
