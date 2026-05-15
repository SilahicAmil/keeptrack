package services

import (
	"os"
	"time"
)

type Logger struct {
	file *os.File
}

type LogEntry struct {
	id        string // unique uuid Ex: 6ba7b810-9dad-11d1-80b4-00c04fd430c8
	function  string
	error     any // error message - I think any is fine for now
	level     string
	timestamp time.Time // date time log happened. based off local user settings
}

func (l *LogEntry) formatLog() {
	// log to be formated into the following string
	// [id] - [timestamp] - [level] - <error>
}

func (l *LogEntry) WriteLog(log LogEntry) {
	// append to or create txt file
	// same location as DB
}

func (l *LogEntry) Error(err any, functionStr string) {

	// Maybe make the id determenistic off the func?
	// So we have consistent id for each functionStr
	// easier to track each function?

	// Pass the error into formatLog
	// Log will be roughly
	// [id] - [timestamp] - <error>
	// <error> will be the raw error from the app
	// Ex: [6ba7b810-9dad-11d1-80b4-00c04fd430c8:CheckAppState] - [05/11/26 03:04:55 PM] - sql: no rows in result set
}
