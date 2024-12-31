package monlog

import (
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

type Logger struct {
	prefix string
	level  Level
	out    io.Writer
}

func NewLogger() *Logger {
	return &Logger{
		prefix: "monkey-logger",
		level:  Debug,
		out:    os.Stdout,
	}
}

func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
	atomic.StoreInt32((*int32)(&l.level), int32(level))
}

func (l *Logger) Info(message string) {
	timestamp := time.Now().Format(time.RFC3339)
	log := fmt.Sprintf("%s [%s] %s: %s", l.prefix, timestamp, "INFO", message)
	fmt.Fprintln(l.out, log)
}

func (l *Logger) Debug(message string) {
	if l.level != Debug {
		return
	}
	timestamp := time.Now().Format(time.RFC3339)
	log := fmt.Sprintf("%s [%s] %s: %s", l.prefix, timestamp, "DEBUG", message)
	fmt.Fprintln(l.out, log)
}

func (l *Logger) Error(message string, err error) {
	timestamp := time.Now().Format(time.RFC3339)
	log := fmt.Sprintf("%s [%s] %s: %s", l.prefix, timestamp, "ERROR", message)
	fmt.Fprintln(l.out, log)
}
