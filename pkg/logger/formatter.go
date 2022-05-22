package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

// DefaultTextFormatter returns a default formatter
func DefaultTextFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339, // "2006-01-02T15:04:05Z07:00"
	}
}

var (
	logstashFields = logrus.Fields{
		"@version": "1",
		"type":     "logger",
	}

	logstashFieldMap = logrus.FieldMap{
		logrus.FieldKeyTime: "@timestamp",
		logrus.FieldKeyMsg:  "message",
	}
)

// NewLogstashFormatter returns a default Logstash formatter:
// A JSON format with "@version" set to "1" (unless set differently in `fields`,
// "type" to "logger" (unless set differently in `fields`),
// "@timestamp" to the logger time and "message" to the logger message.
//
// Note: to set a different configuration use the `LogstashFormatter` structure.
func NewLogstashFormatter(fields logrus.Fields) *logrus.JSONFormatter {
	for k, v := range logstashFields {
		if _, ok := fields[k]; !ok {
			fields[k] = v
		}
	}

	return &logrus.JSONFormatter{FieldMap: logstashFieldMap}
}

func DefaultLogstashFormatter() *logrus.JSONFormatter {
	return NewLogstashFormatter(make(logrus.Fields))
}
