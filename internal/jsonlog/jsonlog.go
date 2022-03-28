package jsonlog

import (
	"encoding/json"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// Define a level type to represent log level

type Level int8

// Initialize constants which represents a specific severity level
// use the iota keywork to assign succesive integer values to the constant
const (
	LevelInfo Level = iota
	LevelError
	LevelFatal
	LevelOff
)

// return a human-friendly string for the severity level
func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

// Define a custom logger type. The holds the output destination that the log entries will be written to
// the minumum log severity level that the log will be written for
type Logger struct {
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

// Return a new Logger Instance which writes log entries
func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
	}
}

// helper methods for printing log types
func (l *Logger) PrintInfo(message string, properties map[string]string) {
	l.print(LevelInfo, message, properties)
}

// Fatal Level
func (l *Logger) Printfatal(err error, properties map[string]string) {
	l.print(LevelError, err.Error(), properties)
	os.Exit(1)
}

func (l *Logger) print(level Level, message string, properties map[string]string) (int, error) {
	if level < l.minLevel {
		return 0, nil
	}

	// Declare anonymous struct holding data for log entry
	aux := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:message`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Level:      level.String(),
		Time:       time.Now().UTC().Format(time.RFC3339),
		Message:    message,
		Properties: properties,
	}
	// Include a stack trace for Error and fatal levels
	if level >= LevelError {
		aux.Trace = string(debug.Stack())
	}

	// declare a line for holding the log entry text
	var line []byte

	// marshall the anonymous struct to JSON
	line, err := json.Marshal(aux)

	if err != nil {
		line = []byte(LevelError.String() + ": Unable to marshal log message:" + err.Error())
	}

	// lock the mutext to prevent concurrent writes to the output destination
	l.mu.Lock()
	defer l.mu.Unlock()

	// Write the log entry followed by a new line
	return l.out.Write(append(line, '\n'))
}

func (l *Logger) Write(message []byte) (n int, err error) {
	return l.print(LevelError, string(message), nil)
}
