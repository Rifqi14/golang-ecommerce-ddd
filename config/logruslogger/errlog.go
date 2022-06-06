package logruslogger

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// TOPIC for setting topic of
	TOPIC = "logging"
	// Error level
	ErrorLevel = log.ErrorLevel
	// Warn level
	WarnLevel = log.WarnLevel
	// Info level
	InfoLevel = log.InfoLevel
)

func LogContext(c string, s string, cor ...interface{}) *log.Entry {
	topic, ok := os.LookupEnv("LOG_TOPIC")
	if !ok {
		topic = TOPIC
	}

	entry := log.WithFields(log.Fields{
		"topic":   topic,
		"context": c,
		"scope":   s,
		"tz":      time.Now().UTC().Format("2006-01-02T15:04:05.999999-07:00"),
	})
	if len(cor) > 0 {
		if cor[0] != nil {
			entry = entry.WithFields(log.Fields{
				"req_id": fmt.Sprintf("%+v", cor[0]),
			})
		}
	}
	return entry
}

func Log(level log.Level, message, context, scope string, corr ...interface{}) {
	log.SetFormatter(&log.JSONFormatter{})
	var correlation interface{}
	if len(corr) > 0 {
		correlation = corr[0]
	}
	entry := LogContext(context, scope, correlation)
	switch level {
	case log.DebugLevel:
		entry.Debug(message)
	case log.InfoLevel:
		entry.Info(message)
	case log.WarnLevel:
		entry.Warn(message)
	case log.ErrorLevel:
		entry.Error(message)
	case log.FatalLevel:
		entry.Fatal(message)
	case log.PanicLevel:
		entry.Panic(message)
	}
}
