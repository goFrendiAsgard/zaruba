package output

import (
	"encoding/csv"
	"os"
	"sync"
)

type RecordLogger interface {
	Log(data ...string) (err error)
}

// CSVRecordLogger is a thread safe csv writer
type CSVRecordLogger struct {
	Mutex    *sync.Mutex
	FileName string
}

// NewCSVRecordLogger create new CSV Log Writer
func NewCSVRecordLogger(fileName string) (c *CSVRecordLogger) {
	c = &CSVRecordLogger{
		Mutex:    &sync.Mutex{},
		FileName: fileName,
	}
	return c
}

// Log will log array of string
func (c *CSVRecordLogger) Log(record ...string) (err error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	f, err := os.OpenFile(c.FileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()
	return writer.Write(record)
}
