package output

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type RecordLogger interface {
	Log(data ...string) (err error)
}

// CSVRecordLogger is a thread safe csv writer
type CSVRecordLogger struct {
	Mutex                  *sync.Mutex
	FileName               string
	logDir                 string
	backupFileNameTemplate string
}

// NewCSVRecordLogger create new CSV Log Writer
func NewCSVRecordLogger(fileName, backupFileNameTemplate string) (c *CSVRecordLogger) {
	absFileName, _ := filepath.Abs(fileName)
	logDir := filepath.Dir(absFileName)
	c = &CSVRecordLogger{
		Mutex:                  &sync.Mutex{},
		FileName:               absFileName,
		logDir:                 logDir,
		backupFileNameTemplate: backupFileNameTemplate,
	}
	return c
}

// Log will log array of string
func (c *CSVRecordLogger) Log(record ...string) (err error) {
	os.Mkdir(c.logDir, os.ModePerm)
	c.Mutex.Lock()
	logFile, err := os.OpenFile(c.FileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		c.Mutex.Unlock()
		return err
	}
	defer logFile.Close()
	logFileStat, err := logFile.Stat()
	if err != nil {
		c.Mutex.Unlock()
		return err
	}
	fileSize := logFileStat.Size()
	if fileSize >= 5*1024*1024 {
		now := time.Now()
		backupFileName := fmt.Sprintf(c.backupFileNameTemplate, now.Format("2006-01-02-15-04-05"))
		absBackupFileName := filepath.Join(c.logDir, backupFileName)
		if err := os.Rename(c.FileName, absBackupFileName); err != nil {
			return err
		}
		c.Mutex.Unlock()
		return c.Log(record...)
	}
	defer c.Mutex.Unlock()
	writer := csv.NewWriter(logFile)
	defer writer.Flush()
	return writer.Write(record)
}
