package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type CustomTextFormatter struct{}

func (f *CustomTextFormatter) Format(entry *log.Entry) ([]byte, error) {
	// Get the file and line number where the log was called
	_, filename, line, _ := runtime.Caller(7)

	// Get the script name from the full file path
	scriptName := filepath.Base(filename)

	// Format the log message
	message := fmt.Sprintf("[%s] [%s] [%s:%d] %s\n",
		entry.Time.Format("2006-01-02 15:04:05"), // Date-time
		entry.Level.String(),                     // Log level
		scriptName,                               // Script name
		line,                                     // Line number
		entry.Message,                            // Log message
	)

	return []byte(message), nil // [2023-08-23 14:58:28] [info] [room.go:66] Room ID: pgb6w37rx1
}

// Logrus configuration for JSON format
func CustomLogConfig() {
	log.SetReportCaller(true)
	formatter := &log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "timestamp",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "message",
		},

		TimestampFormat: "02-01-2006 15:04:05", // the "time" field configuratiom
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			pathname := strings.Split(f.File, "/")
			return "", fmt.Sprintf("%s:%d", pathname[len(pathname)-1], f.Line)
		},
	}
	log.SetFormatter(formatter) // {"file":"exported.go:109","level":"info","message":"Room ID: 57zzb3z13g","timestamp":"23-08-2023 15:01:41"}
}
