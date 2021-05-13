package output

import "sync"

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
