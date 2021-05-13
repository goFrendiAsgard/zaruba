package output

import (
	"fmt"
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

type MockLogger struct {
	Data  []MockLoggerData
	Mutex *sync.Mutex
}

func NewMockLogger() *MockLogger {
	return &MockLogger{
		Data:  []MockLoggerData{},
		Mutex: &sync.Mutex{},
	}
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
	defer m.Mutex.Unlock()
	data := NewMockLoggerData(args...)
	m.Data = append(m.Data, data)
	return 0, nil
}
