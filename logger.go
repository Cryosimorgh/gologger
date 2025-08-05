// logger-client/logger.go
package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Logger struct {
	serviceName string
	loggerURL   string
}

func New(serviceName, loggerURL string) *Logger {
	return &Logger{
		serviceName: serviceName,
		loggerURL:   loggerURL,
	}
}

func (l *Logger) Log(level, message string) error {
	entry := LogEntry{
		Service:   l.serviceName,
		Level:     level,
		Message:   message,
		Timestamp: time.Now(),
	}

	jsonData, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("error marshaling log entry: %v", err)
	}

	resp, err := http.Post(l.loggerURL+"/log", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending log: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (l *Logger) Info(message string) {
	l.Log("INFO", message)
}

func (l *Logger) Error(message string) {
	l.Log("ERROR", message)
}

func (l *Logger) Warn(message string) {
	l.Log("WARN", message)
}

type LogEntry struct {
	Service   string    `json:"service"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
