package logger

import (
	"encoding/csv"
	"os"
	"sync"
	"time"
)

// CSVLogWriter is a thread safe csv writer
type CSVLogWriter struct {
	Mutex    *sync.Mutex
	FileName string
}

// NewCSVLogWriter create new CSV Log Writer
func NewCSVLogWriter(fileName string) (c *CSVLogWriter) {
	c = &CSVLogWriter{
		Mutex:    &sync.Mutex{},
		FileName: fileName,
	}
	return c
}

// Log will log array of string
func (c *CSVLogWriter) Log(data ...string) (err error) {
	record := []string{time.Now().String()}
	record = append(record, data...)
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	f, err := os.OpenFile(c.FileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	defer f.Close()
	if err != nil {
		return err
	}
	writer := csv.NewWriter(f)
	defer writer.Flush()
	return writer.Write(record)
}
