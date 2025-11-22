package logger

import (
	"log"
	"os"
)

type StdLogger struct {
	logger *log.Logger
}

func NewStdLogger() Logger {
	return &StdLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

func (l *StdLogger) Debug(msg string, args ...interface{}) {
	l.logger.Printf("[DEBUG] "+msg, args...)
}

func (l *StdLogger) Info(msg string, args ...interface{}) {
	l.logger.Printf("[INFO] "+msg, args...)
}

func (l *StdLogger) Warn(msg string, args ...interface{}) {
	l.logger.Printf("[WARN] "+msg, args...)
}

func (l *StdLogger) Error(msg string, args ...interface{}) {
	l.logger.Printf("[ERROR] "+msg, args...)
}

func (l *StdLogger) Fatal(msg string, args ...interface{}) {
	l.logger.Fatalf("[FATAL] "+msg, args...)
}
