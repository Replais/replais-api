package logger

// Logger defines the interface for logging operations.
// This allows you to swap implementations (std log, zap, logrus, etc.) easily.
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}
