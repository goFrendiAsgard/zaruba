package output

import (
	"fmt"
	"sync"
)

type MockRecordLogger struct {
	Mutex *sync.Mutex
	Data  [][]string
}

func NewMockRecordLogger() *MockRecordLogger {
	return &MockRecordLogger{
		Mutex: &sync.Mutex{},
		Data:  [][]string{},
	}
}

func (m *MockRecordLogger) Log(data ...string) (err error) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.Data = append(m.Data, data)
	return nil
}

type MockInvalidRecordLogger struct{}

func NewMockInvalidRecordLogger() *MockInvalidRecordLogger {
	return &MockInvalidRecordLogger{}
}

func (mi *MockInvalidRecordLogger) Log(data ...string) (err error) {
	return fmt.Errorf("cannot write")
}
