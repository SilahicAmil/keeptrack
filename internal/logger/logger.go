package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Logger struct {
	file *os.File
}

type LogEntry struct {
	id        string // unique uuid Ex: 6ba7b810-9dad-11d1-80b4-00c04fd430c8
	function  string
	error     any // error message - I think any is fine for now
	level     string
	timestamp string // date time log happened. based off local user settings
}

func (l *LogEntry) formatLog() string {
	// log to be formated into the following string
	// [id] - [timestamp] - [level] - <error>
	return fmt.Sprintf("[%s] - [%s] - [%s] - %s",
		l.id, l.timestamp, l.level, l.error)
}

func (l *LogEntry) WriteLog() error {

	configDir, err := os.UserConfigDir()

	if err != nil {
		return err
	}

	appDir := filepath.Join(configDir, "keeptrack")

	if err := os.MkdirAll(appDir, os.ModePerm); err != nil {
		return err
	}

	logPath := filepath.Join(appDir, "keeptrack.log")

	f, err := os.OpenFile(
		logPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(
		l.formatLog() + "\n\n",
	)

	return err
}

func Error(err any, functionStr string) {
	// set struct
	errLog := &LogEntry{
		id:        uuid.New().String(),
		function:  functionStr,
		error:     err,
		level:     "error",
		timestamp: time.Now().Format("01-02-2006 15:04:05"),
	}

	// frmt struct to string
	errLog.formatLog()

	// write log
	errLog.WriteLog()

	// Pass the error into formatLog
	// Log will be roughly
	// [id] - [timestamp] - <error>
	// <error> will be the raw error from the app
	// Ex: [6ba7b810-9dad-11d1-80b4-00c04fd430c8:CheckAppState] - [05/11/26 03:04:55 PM] - sql: no rows in result set
}
